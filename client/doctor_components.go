package client

import (
	"image/color"
	"log"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dhanushs3366/patient-app/db/models"
)

func (c *Client) DoctorLogin() *fyne.Container {
	mailEntry := widget.NewEntry()
	mailEntry.SetPlaceHolder("Enter email")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	loginButton := widget.NewButton("Login", func() {
		doctor, err := c.store.GetDoctorByMail(mailEntry.Text)

		if err != nil {
			c.showCustomPrompt("Warning!", "Invalid Credentials")
		} else {
			if doctor.Password != passwordEntry.Text {
				c.showCustomPrompt("Warning!", "Invalid Credentials")
			}

			c.loggedInDoctor = doctor
			log.Printf("%v\n", doctor)
			// reroute to their list of appointments
		}
	})

	loginForm := container.New(
		layout.NewGridLayoutWithRows(3),
		layout.NewSpacer(),
		container.New(
			layout.NewGridLayoutWithColumns(3),
			layout.NewSpacer(),
			container.NewVBox(
				mailEntry,
				passwordEntry,
				loginButton,
			),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	// Set button importance to change its color

	// Create layout with form and button

	return loginForm
}

func (c *Client) BookDoctor() *fyne.Container {
	doctorCardSize := fyne.NewSize(400, 250)

	var wg sync.WaitGroup
	var doctors []models.Doctor
	var filteredDoctors []models.Doctor
	var specs []string //used to store all different type of specialisations

	wg.Add(1)
	go func() {
		defer wg.Done()
		doctorsDB, err := c.store.GetAllDoctors()
		if err != nil {
			log.Printf("Err %s\n", err.Error())
		}
		doctors = append(doctors, doctorsDB...)
		filteredDoctors = make([]models.Doctor, len(doctorsDB))
		dbSpecs, err := c.store.GetAllSpecialisationTypes()
		if err != nil {
			log.Printf("Error %s\n", err.Error())
		}
		specs = append(specs, dbSpecs...)
	}()

	wg.Wait()
	copy(filteredDoctors, doctors)
	doctorCards := c.makeCards(filteredDoctors)
	doctorContainer := container.New(layout.NewGridWrapLayout(doctorCardSize), doctorCards...)

	// use Select.Focus to add effects when you hover over it
	selectBtnHandler := func(option string) {
		copy(filteredDoctors, doctors)
		for i, doctor := range filteredDoctors {
			if doctor.Specialisation.Name != option {
				filteredDoctors = append(filteredDoctors[:i], filteredDoctors[i:]...)
			}
		}
		doctorContainer.Refresh()
	}

	selectBtn := widget.NewSelect(specs, selectBtnHandler)
	dropDownContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), selectBtn)

	doctorContainer = container.NewBorder(dropDownContainer, nil, nil, nil, doctorContainer)
	return doctorContainer
}

func (c *Client) makeCards(doctors []models.Doctor) []fyne.CanvasObject {
	var cards []fyne.CanvasObject
	for _, doctor := range doctors {

		log.Printf("%v", doctor.Specialisation)

		specialtyText := canvas.NewText(doctor.Specialisation.Name, theme.Color(theme.ColorNameForeground))

		descriptionLabel := widget.NewLabel(doctor.Specialisation.Description)
		descriptionLabel.Wrapping = fyne.TextWrapWord

		contactContainer := container.NewHBox(
			widget.NewIcon(theme.MailComposeIcon()),
			widget.NewLabel(doctor.Email),
			widget.NewIcon(theme.MailReplyIcon()),
			widget.NewLabel(doctor.Phone),
		)

		// book appointment takes in a doctor and redirects to a container that takes patient form and along with doctor
		// and then on submit we create appointment
		bookAppointmentHandler := func(doctor models.Doctor) func() {
			return func() {
				nameLabel := widget.NewLabel("Name:")
				nameEntry := widget.NewEntry()

				emailLabel := widget.NewLabel("Email:")
				emailEntry := widget.NewEntry()

				phoneLabel := widget.NewLabel("Phone Number:")
				phoneEntry := widget.NewEntry()

				entryMap := map[*widget.Label]*widget.Entry{
					nameLabel:  nameEntry,
					emailLabel: emailEntry,
					phoneLabel: phoneEntry,
				}
				submit := widget.NewButton("Submit", func() {
					isValid, invalidStr := checkValidForms("Booking", "Booking is done succesfully", &c.Window, entryMap)

					if !isValid {
						promptWindow("Booking", invalidStr, &c.Window)
						return
					}

					// make a db transaction
					// check if patient alr exists
					patient, err := c.store.GetPatientByEmail(emailEntry.Text)
					log.Printf("\n\npatientID: %v\n", patient)
					if err != nil {
						// new patient
						patientID, err := c.store.CreateNewPatient(models.Patient{
							Name:  nameEntry.Text,
							Email: emailEntry.Text,
							Phone: phoneEntry.Text,
						})

						log.Printf("\n\npatientID: %d\n", patientID)
						if err != nil {
							// only for testing
							promptWindow("Inserting Patient", err.Error(), &c.Window)
							return
						}

						// book appointment for new patient
						err = c.store.BookAppointment(patientID, doctor.ID)

						if err != nil {
							promptWindow("Failed", err.Error(), &c.Window)
							return
						}

						promptWindow("Success", "Your appointment has been booked succesfully", &c.Window)
						return
					}

					err = c.store.BookAppointment(patient.ID, doctor.ID)

					if err != nil {
						promptWindow("Failed", err.Error(), &c.Window)
						return
					}
					promptWindow("Success", "Your appointment has been booked succesfully", &c.Window)
					c.Window.SetContent(c.Navbar(c.About()))
				})

				cancel := widget.NewButton("Cancel", func() {
					c.Window.SetContent(c.Navbar(c.BookDoctor()))
				})

				patientForm := container.New(
					layout.NewFormLayout(),
					nameLabel, nameEntry,
					emailLabel, emailEntry,
					phoneLabel, phoneEntry,
					widget.NewLabel(""), container.New(layout.NewGridLayoutWithColumns(2), submit, cancel),
				)

				c.Window.SetContent(c.Navbar(patientForm))

			}
		}

		bookButton := widget.NewButton("Book", bookAppointmentHandler(doctor))

		var availabilityText *canvas.Text
		if doctor.IsAvailable {
			availabilityText = canvas.NewText("Available", color.NRGBA{R: 0, G: 180, B: 0, A: 255}) // Green
		} else {
			availabilityText = canvas.NewText("Not Available", color.NRGBA{R: 180, G: 0, B: 0, A: 255}) // Red
		}
		availabilityText.Alignment = fyne.TextAlignCenter
		availabilityText.TextStyle = fyne.TextStyle{Bold: true}

		cardContent := container.NewVBox(
			container.New(layout.NewCustomPaddedHBoxLayout(10), specialtyText),
			descriptionLabel,
			contactContainer,
			container.New(layout.NewGridLayoutWithColumns(2), bookButton, availabilityText),
		)
		// 403d39

		card := widget.NewCard(doctor.Name, "", cardContent)

		cards = append(cards, card)
	}
	return cards
}
