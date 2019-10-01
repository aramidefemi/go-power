package config
// Config contains those necessary fields that Mailer needs to send e-mails.
type Config struct {
    // Host is the server mail host, IP or address.
    Host string
    // Port is the listening port.
    Port int
    // Username is the auth username@domain.com for the sender.
    Username string
    // Password is the auth password for the sender.
    Password string
    // FromAddr is the 'from' part of the mail header, it overrides the username.
    FromAddr string
    // FromAlias is the from part, if empty this is the first part before @ from the Username field.
    FromAlias string
    // UseCommand enable it if you want to send e-mail with the mail command  instead of smtp.
    //
    // Host,Port & Password will be ignored.
    // ONLY FOR UNIX.
    UseCommand bool
}

var MailConfig Config = {
	Host:     "smtp.mailgun.org",
	Username: "postmaster",
	Password: "38304272b8ee5c176d5961dc155b2417",
	FromAddr: "postmaster@sandbox661c307650f04e909150b37c0f3b2f09.mailgun.org",
	Port:     587,
	// Enable UseCommand to support sendmail unix command,
	// if this field is true then Host, Username, Password and Port are not required,
	// because these info already exists in your local sendmail configuration.
	//
	// Defaults to false.
	UseCommand: false,
}