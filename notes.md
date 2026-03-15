conferenceName := "Go Conference"
const numberConferenceTicketsAvailable = 50
var remainingConferenceTickets uint = 50

/ /var bookings = [50]string{"Sascha", "Sasc", "Sas", "Sa", "S"}
//var bookings = [50]string{}
//var bookings [50]string => array
//var bookings = []string{} //=> slice
//bookings := []string{} //=> slice

Create map
var userData = make(map[string]string)
var mymap map[string]string
userData["firstName"] = firstName
userData["lastName"] = lastName
userData["email"] = email
fmt.Printf("User data is %v\n", userData)
=> no mixed data types allowed in map, all values must be of the same type
userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)
list of maps => var bookings = make([]map[string]string, 0)
