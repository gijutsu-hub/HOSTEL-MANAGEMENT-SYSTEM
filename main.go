package main

import (
	"github.com/go-gomail/gomail"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Userdetails struct {
	Regno         string `gorm:"not null;unique"`
	Roomid        string
	Roomtype      int8
	Block         int8
	Stu           int8
	Price         float64
	Dateofbooking time.Time
}

type Admin struct {
	Username string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
}
type User struct {
	Regno    string    `gorm:"primaryKey"`
	Email    string    `gorm:"unique"`
	Password string    `gorm:"not null"`
	Startime time.Time `gorm:"default:not null"`
	Endtime  time.Time `gorm:"default:not null"`
}
type Room struct {
	Roomid   string `gorm:"primaryKey"`
	Roomtype int8
	Block    int8
	Stu      int8
	Price    float64
}
type Sturoomdet struct {
	Regno  string `gorm:"not null;unique"`
	Roomid string `gorm:"not null;unique"`
}

func GeneratePassword(length int) string {
	var password string
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time
	for i := 0; i < length; i++ {
		// Generate a random integer between 0 and 62 (inclusive)
		r := rand.Intn(62)
		// Convert the integer to a character
		if r < 26 {
			password += string('a' + rune(r))
		} else if r < 52 {
			password += string('A' + rune(r-26))
		} else {
			password += string('0' + rune(r-52))
		}
	}
	return password
}

func sendemail1(email string, username string, password string) bool {

	to := []string{email}
	subject := "User Created Success"
	body := `
	<!doctype html>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>Booking Template</title>
    <style>
@media only screen and (max-width: 620px) {
  table.body h1 {
    font-size: 28px !important;
    margin-bottom: 10px !important;
  }

  table.body p,
table.body ul,
table.body ol,
table.body td,
table.body span,
table.body a {
    font-size: 16px !important;
  }

  table.body .wrapper,
table.body .article {
    padding: 10px !important;
  }

  table.body .content {
    padding: 0 !important;
  }

  table.body .container {
    padding: 0 !important;
    width: 100% !important;
  }

  table.body .main {
    border-left-width: 0 !important;
    border-radius: 0 !important;
    border-right-width: 0 !important;
  }

  table.body .btn table {
    width: 100% !important;
  }

  table.body .btn a {
    width: 100% !important;
  }

  table.body .img-responsive {
    height: auto !important;
    max-width: 100% !important;
    width: auto !important;
  }
}
@media all {
  .ExternalClass {
    width: 100%;
  }

  .ExternalClass,
.ExternalClass p,
.ExternalClass span,
.ExternalClass font,
.ExternalClass td,
.ExternalClass div {
    line-height: 100%;
  }

  .apple-link a {
    color: inherit !important;
    font-family: inherit !important;
    font-size: inherit !important;
    font-weight: inherit !important;
    line-height: inherit !important;
    text-decoration: none !important;
  }

  #MessageViewBody a {
    color: inherit;
    text-decoration: none;
    font-size: inherit;
    font-family: inherit;
    font-weight: inherit;
    line-height: inherit;
  }

  .btn-primary table td:hover {
    background-color: #34495e !important;
  }

  .btn-primary a:hover {
    background-color: #34495e !important;
    border-color: #34495e !important;
  }
}
</style>
  </head>
  <body style="background-color: #f6f6f6; font-family: sans-serif; -webkit-font-smoothing: antialiased; font-size: 14px; line-height: 1.4; margin: 0; padding: 0; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;">
    <span class="preheader" style="color: transparent; display: none; height: 0; max-height: 0; max-width: 0; opacity: 0; overflow: hidden; mso-hide: all; visibility: hidden; width: 0;">Booking conformation for hostel</span>
    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="body" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #f6f6f6; width: 100%;" width="100%" bgcolor="#f6f6f6">
      <tr>
        <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;" valign="top">&nbsp;</td>
        <td class="container" style="font-family: sans-serif; font-size: 14px; vertical-align: top; display: block; max-width: 580px; padding: 10px; width: 580px; margin: 0 auto;" width="580" valign="top">
          <div class="content" style="box-sizing: border-box; display: block; margin: 0 auto; max-width: 580px; padding: 10px;">

            <!-- START CENTERED WHITE CONTAINER -->
            <table role="presentation" class="main" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background: #ffffff; border-radius: 3px; width: 100%;" width="100%">

              <!-- START MAIN CONTENT AREA -->
              <tr>
                <td class="wrapper" style="font-family: sans-serif; font-size: 14px; vertical-align: top; box-sizing: border-box; padding: 20px;" valign="top">
                  <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                    <tr>
                      <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;" valign="top">
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">Hi ` + username + `,</p>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">Your user have been creadted successfully the mentions login password is <b>` + password + `</b></p>
                        <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="btn btn-primary" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; box-sizing: border-box; width: 100%;" width="100%">
                        </table>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">This is a system generated email, please do not reply back to it. If you are facing any issue try to contact chief warden by emailing at <a href="mailto:cw@vitbhopal.ac.in">cw@vitbhopal.ac.in</a></p><br><br>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">Best Wishes</p>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">VITB HRS BOT</p>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>

            <!-- END MAIN CONTENT AREA -->
            </table>
            <!-- END CENTERED WHITE CONTAINER -->

            <!-- START FOOTER -->
            <div class="footer" style="clear: both; margin-top: 10px; text-align: center; width: 100%;">
              <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                <tr>
                  <td class="content-block" style="font-family: sans-serif; vertical-align: top; padding-bottom: 10px; padding-top: 10px; color: #999999; font-size: 12px; text-align: center;" valign="top" align="center">
                    <span class="apple-link" style="color: #999999; font-size: 12px; text-align: center;">VIT BHOPAL UNIVERSITY,<br> Bhopal-Indore Highway Kothrikalan, Sehore<br>  Madhya Pradesh – 466114
                    </span>
                  </td>
                </tr>
                <tr>
                  <td class="content-block powered-by" style="font-family: sans-serif; vertical-align: top; padding-bottom: 10px; padding-top: 10px; color: #999999; font-size: 12px; text-align: center;" valign="top" align="center">
                    Powered by <a href="#" style="color: #999999; font-size: 12px; text-align: center; text-decoration: none;">VITBHOPAL</a>.
                  </td>
                </tr>
              </table>
            </div>
            <!-- END FOOTER -->

          </div>
        </td>
        <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;" valign="top">&nbsp;</td>
      </tr>
    </table>
  </body>
</html>
	`

	// Replace dynamic values in the email template

	// Define mailer
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_USER"))
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	// Define SMTP server details
	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), 465, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	// Send email via SMTP server
	if err := dialer.DialAndSend(mailer); err != nil {
		return false
	} else {
		return true
	}

}

func sendemail(email string, username string) bool {

	to := []string{email}
	subject := "Hostel Allotment success"
	body := `
	<!doctype html>
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>Booking Template</title>
    <style>
@media only screen and (max-width: 620px) {
  table.body h1 {
    font-size: 28px !important;
    margin-bottom: 10px !important;
  }

  table.body p,
table.body ul,
table.body ol,
table.body td,
table.body span,
table.body a {
    font-size: 16px !important;
  }

  table.body .wrapper,
table.body .article {
    padding: 10px !important;
  }

  table.body .content {
    padding: 0 !important;
  }

  table.body .container {
    padding: 0 !important;
    width: 100% !important;
  }

  table.body .main {
    border-left-width: 0 !important;
    border-radius: 0 !important;
    border-right-width: 0 !important;
  }

  table.body .btn table {
    width: 100% !important;
  }

  table.body .btn a {
    width: 100% !important;
  }

  table.body .img-responsive {
    height: auto !important;
    max-width: 100% !important;
    width: auto !important;
  }
}
@media all {
  .ExternalClass {
    width: 100%;
  }

  .ExternalClass,
.ExternalClass p,
.ExternalClass span,
.ExternalClass font,
.ExternalClass td,
.ExternalClass div {
    line-height: 100%;
  }

  .apple-link a {
    color: inherit !important;
    font-family: inherit !important;
    font-size: inherit !important;
    font-weight: inherit !important;
    line-height: inherit !important;
    text-decoration: none !important;
  }

  #MessageViewBody a {
    color: inherit;
    text-decoration: none;
    font-size: inherit;
    font-family: inherit;
    font-weight: inherit;
    line-height: inherit;
  }

  .btn-primary table td:hover {
    background-color: #34495e !important;
  }

  .btn-primary a:hover {
    background-color: #34495e !important;
    border-color: #34495e !important;
  }
}
</style>
  </head>
  <body style="background-color: #f6f6f6; font-family: sans-serif; -webkit-font-smoothing: antialiased; font-size: 14px; line-height: 1.4; margin: 0; padding: 0; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;">
    <span class="preheader" style="color: transparent; display: none; height: 0; max-height: 0; max-width: 0; opacity: 0; overflow: hidden; mso-hide: all; visibility: hidden; width: 0;">Booking conformation for hostel</span>
    <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="body" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #f6f6f6; width: 100%;" width="100%" bgcolor="#f6f6f6">
      <tr>
        <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;" valign="top">&nbsp;</td>
        <td class="container" style="font-family: sans-serif; font-size: 14px; vertical-align: top; display: block; max-width: 580px; padding: 10px; width: 580px; margin: 0 auto;" width="580" valign="top">
          <div class="content" style="box-sizing: border-box; display: block; margin: 0 auto; max-width: 580px; padding: 10px;">

            <!-- START CENTERED WHITE CONTAINER -->
            <table role="presentation" class="main" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background: #ffffff; border-radius: 3px; width: 100%;" width="100%">

              <!-- START MAIN CONTENT AREA -->
              <tr>
                <td class="wrapper" style="font-family: sans-serif; font-size: 14px; vertical-align: top; box-sizing: border-box; padding: 20px;" valign="top">
                  <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                    <tr>
                      <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;" valign="top">
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">Hi ` + username + `,</p>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">You have booked your room and your invoice will be generated on vtop make. This is just a conformation email.</p>
                        <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="btn btn-primary" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; box-sizing: border-box; width: 100%;" width="100%">
                        </table>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">This is a system generated email, please do not reply back to it. If you are facing any issue try to contact chief warden by emailing at <a href="mailto:cw@vitbhopal.ac.in">cw@vitbhopal.ac.in</a></p><br><br>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">Best Wishes</p>
                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; margin-bottom: 15px;">VITB HRS BOT</p>
                      </td>
                    </tr>
                  </table>
                </td>
              </tr>

            <!-- END MAIN CONTENT AREA -->
            </table>
            <!-- END CENTERED WHITE CONTAINER -->

            <!-- START FOOTER -->
            <div class="footer" style="clear: both; margin-top: 10px; text-align: center; width: 100%;">
              <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                <tr>
                  <td class="content-block" style="font-family: sans-serif; vertical-align: top; padding-bottom: 10px; padding-top: 10px; color: #999999; font-size: 12px; text-align: center;" valign="top" align="center">
                    <span class="apple-link" style="color: #999999; font-size: 12px; text-align: center;">VIT BHOPAL UNIVERSITY,<br> Bhopal-Indore Highway Kothrikalan, Sehore<br>  Madhya Pradesh – 466114
                    </span>
                  </td>
                </tr>
                <tr>
                  <td class="content-block powered-by" style="font-family: sans-serif; vertical-align: top; padding-bottom: 10px; padding-top: 10px; color: #999999; font-size: 12px; text-align: center;" valign="top" align="center">
                    Powered by <a href="#" style="color: #999999; font-size: 12px; text-align: center; text-decoration: none;">VITBHOPAL</a>.
                  </td>
                </tr>
              </table>
            </div>
            <!-- END FOOTER -->

          </div>
        </td>
        <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;" valign="top">&nbsp;</td>
      </tr>
    </table>
  </body>
</html>
	`

	// Define mailer
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("SMTP_USER"))
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	// Define SMTP server details
	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), 465, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	// Send email via SMTP server
	if err := dialer.DialAndSend(mailer); err != nil {
		return false
	} else {
		return true
	}

}

var env = godotenv.Load()
var dsn = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
var store = sessions.NewCookieStore([]byte(os.Getenv("COOKIE_SECRET")))

func HomeHandeler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := template.Must(template.ParseFiles("templates/Home/home.html")).Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginHandeler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login")
	now := time.Now()
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/Home/login.html")).Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		regno := r.FormValue("regno")
		pass := r.FormValue("password")
		var user User
		result := db.Where("regno = ? AND startime < ?", regno, now).First(&user)
		if result.Error != nil {
			http.Redirect(w, r, "/login?msg=Your registration time haven't started", http.StatusSeeOther)
			return
		} else {
			result := db.Where("regno = ? AND endtime > ?", regno, now).First(&user)
			if result.Error != nil {
				http.Redirect(w, r, "/login?msg=Your registration is closed", http.StatusSeeOther)
				return
			} else {
				result := db.Where("regno = ? AND password = ?", regno, pass).First(&user)
				if result.Error != nil {
					http.Redirect(w, r, "/login?msg=Bad+credentials", http.StatusSeeOther)
					return
				} else {
					session.Values["username"] = regno
					session.Save(r, w)
					http.Redirect(w, r, "/hostel-registration", http.StatusSeeOther)
					return
				}
			}
		}
	}

}

func checkhostel(username string) bool {
	query := "SELECT COUNT(*) FROM sturoomdets WHERE regno = ? "
	var count int
	db.Raw(query, username).Scan(&count)
	if err != nil {
		return false
	} else {
		if count > 0 {
			return true
		} else {
			return false
		}

	}
}

func hostelreg(regno string, roomid string) bool {
	login := Sturoomdet{Regno: regno, Roomid: roomid}
	result := db.Select("regno", "roomid").Create(&login)
	if result.Error != nil {
		return false
	} else {
		return true
	}
}

func hosteldet(roomid string) {
	sql := "UPDATE roomdeatils SET stu = stu -1 WHERE roomid = ?"
	db.Exec(sql, roomid)
}

func Hostelregister(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "login")
	var rooms []Room
	if err := db.Table("roomdeatils").Select("roomid, roomtype, Block, stu, Price").Scan(&rooms).Error; err != nil {
		panic(err)
	}
	username, ok := session.Values["username"].(string)
	if !ok {
		http.Redirect(w, r, "/login?msg=Session+expired", http.StatusSeeOther)
		return
	} else {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "text/html")
			err := template.Must(template.ParseFiles("templates/Home/hostel.html")).Execute(w, rooms)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		} else {
			roomid := r.FormValue("roomid")
			db.Table("roomdeatils").Where("roomid = ?", roomid).Find(&rooms)
			for _, room := range rooms {
				if room.Stu == room.Roomtype {
					http.Redirect(w, r, "/hostel-registration?msg=Room+is+full+you+cannot+book", http.StatusSeeOther)
					return
				} else if checkhostel(username) {
					http.Redirect(w, r, "/hostel-registration?msg=You+Have+already+booked+the+room", http.StatusSeeOther)
					return
				} else {
					hostelreg(username, roomid)
					hosteldet(roomid)
					var user User
					db.Where("regno = ?", username).First(&user)
					email := user.Email
					sendemail(email, username)
					http.Redirect(w, r, "/hostel-registration?msg=booking+succcess+your+invoice+will+generated+shortly", http.StatusSeeOther)
					return
				}
			}
			http.Redirect(w, r, "/hostel-registration?msg=You+modified+the+room+id+booking+not+conformed", http.StatusSeeOther)
			return

		}
	}
}

func AdminHandeler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "admin")
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/login.html")).Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		regno := r.FormValue("regno")
		pass := r.FormValue("password")
		var admin Admin

		result := db.Where("username = ? AND password = ?", regno, pass).First(&admin)
		if result.Error != nil {
			http.Redirect(w, r, "/hrsadmin?msg=Bad+credentials", http.StatusSeeOther)
			return
		} else {
			session.Values["username"] = regno
			session.Save(r, w)
			http.Redirect(w, r, "/hostel-admin", http.StatusSeeOther)
			return
		}
	}
}
func Hosteladmin(w http.ResponseWriter, r *http.Request) {
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err1 := template.Must(template.ParseFiles("templates/admin/admin.html")).Execute(w, nil)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func ManageroomHandler(w http.ResponseWriter, r *http.Request) {
	var rooms []Room
	if err := db.Table("roomdeatils").Select("roomid, roomtype, Block, stu, Price").Scan(&rooms).Error; err != nil {
		panic(err)
	}
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/mgt/hostel.html")).Execute(w, rooms)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		roomid := r.FormValue("roomid")
		roomprice := r.FormValue("roomprice")
		roomtype := r.FormValue("roomtype")
		block := r.FormValue("block")
		sql := "UPDATE roomdeatils SET roomtype = ?, Block=?, Price=? WHERE roomid = ?"
		db.Exec(sql, roomtype, block, roomprice, roomid)
		http.Redirect(w, r, "/hrsadmin-managerooms?msg=update+successfull", http.StatusSeeOther)
		return
	}
}
func UserloginHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	if err := db.Table("users").Select("regno, email, startime, endtime").Scan(&users).Error; err != nil {
		panic(err)
	}
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/urslg/main.html")).Execute(w, users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		regno := r.FormValue("regno")
		email := r.FormValue("email")
		start := r.FormValue("st")
		et := r.FormValue("et")
		sql := "UPDATE users SET email=?, startime=?, endtime=? WHERE regno = ?"
		db.Exec(sql, email, start, et, regno)
		http.Redirect(w, r, "/hrsadmin-managelogin?msg=update+successfull", http.StatusSeeOther)
		return
	}
}

func AllroomHandler(w http.ResponseWriter, r *http.Request) {
	var rooms []Room
	if err := db.Table("roomdeatils").Select("roomid, roomtype, Block, stu, Price").Scan(&rooms).Error; err != nil {
		panic(err)
	}
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/usrmgt/hostel.html")).Execute(w, rooms)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		username := r.FormValue("regno")
		roomid := r.FormValue("roomid")

		db.Table("roomdeatils").Where("roomid = ?", roomid).Find(&rooms)
		for _, room := range rooms {
			if room.Stu == room.Roomtype {
				http.Redirect(w, r, "/hrsadmin-allot-room?msg=Room+is+full+you+cannot+book", http.StatusSeeOther)
				return
			} else if checkhostel(username) {
				http.Redirect(w, r, "/hrsadmin-allot-room?msg=You+Have+already+booked+the+room", http.StatusSeeOther)
				return
			} else {
				hostelreg(username, roomid)
				hosteldet(roomid)
				var user User
				db.Where("regno = ?", username).First(&user)
				email := user.Email
				sendemail(email, username)
				http.Redirect(w, r, "/hrsadmin-allot-room?msg=booking+succcess+your+invoice+will+generated+shortly", http.StatusSeeOther)
				return
			}
		}
		http.Redirect(w, r, "/hrsadmin-allot-room?msg=You+modified+the+room+id+booking+not+conformed", http.StatusSeeOther)
		return

	}
}

func AddroomHandler(w http.ResponseWriter, r *http.Request) {
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/mgt/add.html")).Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		roomid := r.FormValue("roomid")
		roomprice := r.FormValue("roomprice")
		roomtype := r.FormValue("roomtype")
		block := r.FormValue("block")
		sql := "INSERT INTO `roomdeatils`(roomid, roomtype, Block, Price) VALUES (?,?,?,?)"
		db.Exec(sql, roomid, roomtype, block, roomprice)
		http.Redirect(w, r, "/hrsadmin-addrooms?msg=room+added+successfull", http.StatusSeeOther)
		return
	}
}

func AdduserHandler(w http.ResponseWriter, r *http.Request) {
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/urslg/add.html")).Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		regno := r.FormValue("regno")
		email := r.FormValue("email")
		t := time.Now()
		e := t.Add(6 * time.Hour)
		password := GeneratePassword(10)
		sql := "INSERT INTO users(regno, email, password, startime, endtime) VALUES (?,?,?,?,?)"
		db.Exec(sql, regno, email, password, t, e)
		sendemail1(email, regno, password)
		http.Redirect(w, r, "/adduser?msg=User+created+Successfull", http.StatusSeeOther)
		return
	}
}

func Logouthandeler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login")
	if err != nil {
		http.Redirect(w, r, "/login?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	session.Options.MaxAge = -1 // set MaxAge to -1 to delete the session
	session.Save(r, w)
	http.Redirect(w, r, "/login?msg=logout+successfully", http.StatusSeeOther)
	return
}

func Logoutadminhandeler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	session.Options.MaxAge = -1 // set MaxAge to -1 to delete the session
	session.Save(r, w)
	http.Redirect(w, r, "/hrsadmin?msg=logout+successfully", http.StatusSeeOther)
	return
}

func DeletroomHandler(w http.ResponseWriter, r *http.Request) {
	var rooms []Room
	if err := db.Table("roomdeatils").Select("roomid, roomtype, Block, stu, Price").Scan(&rooms).Error; err != nil {
		panic(err)
	}
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/mgt/delete.html")).Execute(w, rooms)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		roomid := r.FormValue("roomid")
		sql := "DELETE FROM `roomdeatils` WHERE roomid = ?"
		db.Exec(sql, roomid)
		http.Redirect(w, r, "/hrsadmin-delete?msg=deleted+successfull", http.StatusSeeOther)
		return
	}
}

func Userdeletehandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	if err := db.Table("users").Select("regno, email, startime, endtime").Scan(&users).Error; err != nil {
		panic(err)
	}
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/urslg/del.html")).Execute(w, users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		regno := r.FormValue("regno")
		sql := "DELETE FROM `users` WHERE regno = ?"
		db.Exec(sql, regno)
		http.Redirect(w, r, "/delete-user?msg=deleted+successfull", http.StatusSeeOther)
		return
	}
}
func DeletallotedroomHandler(w http.ResponseWriter, r *http.Request) {
	var data []Userdetails
	query := `SELECT * FROM sturoomdets JOIN roomdeatils ON sturoomdets.roomid = roomdeatils.roomid`
	db.Raw(query).Scan(&data)
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/usrmgt/del.html")).Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		roomid := r.FormValue("roomid")
		regno := r.FormValue("regno")
		sql := "DELETE FROM `sturoomdets` WHERE regno = ?"
		db.Exec(sql, regno)
		hosteldet(roomid)
		http.Redirect(w, r, "/hrsadmin-delete-alloted-room?msg=room+deleted+successfully", http.StatusSeeOther)
		return
	}
}

func ManageuserHandler(w http.ResponseWriter, r *http.Request) {
	var data []Userdetails
	query := `SELECT * FROM sturoomdets JOIN roomdeatils ON sturoomdets.roomid = roomdeatils.roomid`
	db.Raw(query).Scan(&data)
	_, err := store.Get(r, "admin")
	if err != nil {
		http.Redirect(w, r, "/hrsadmin?msg=Invalid+session", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")
		err := template.Must(template.ParseFiles("templates/admin/usrmgt/user.html")).Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else if r.Method == "POST" {
		roomid := r.FormValue("roomid")
		sql := "DELETE FROM `roomdeatils` WHERE roomid = ?"
		db.Exec(sql, roomid)
		http.Redirect(w, r, "/hrsadmin-delete?msg=deleted+successfull", http.StatusSeeOther)
		return
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandeler)
	http.HandleFunc("/login", LoginHandeler)
	http.HandleFunc("/hrsadmin-managelogin", UserloginHandler)
	http.HandleFunc("/hrsadmin", AdminHandeler)
	http.HandleFunc("/hostel-registration", Hostelregister)
	http.HandleFunc("/hostel-admin", Hosteladmin)
	http.HandleFunc("/adduser", AdduserHandler)
	//http.HandleFunc("/delete-user", DeletuserHandler)
	http.HandleFunc("/hrsadmin-manageuser", ManageuserHandler)
	http.HandleFunc("/hrsadmin-managerooms", ManageroomHandler)
	http.HandleFunc("/hrsadmin-addrooms", AddroomHandler)
	http.HandleFunc("/hrsadmin-allot-room", AllroomHandler)
	http.HandleFunc("/hrsadmin-delete", DeletroomHandler)
	http.HandleFunc("/hrsadmin-delete-alloted-room", DeletallotedroomHandler)
	http.HandleFunc("/logout-admin", Logoutadminhandeler)
	http.HandleFunc("/delete-user", Userdeletehandler)
	http.HandleFunc("/logout", Logouthandeler)

	if err != nil {
		panic("failed to connect database")
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
