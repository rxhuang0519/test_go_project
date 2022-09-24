package models

type Model interface {
	Base | Message | User
}
