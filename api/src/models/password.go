package models

//Password is used in the UpdatePassword funcion

type Password struct {
	Passwd    string `json:"passwd"`
	Newpasswd string `json:"newpasswd"`
}
