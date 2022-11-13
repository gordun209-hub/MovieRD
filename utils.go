package main

const UserLimit = 5

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
