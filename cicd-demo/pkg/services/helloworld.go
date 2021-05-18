package services


var (
	Version = "v-1.0.32"
)

func Output()(word string, err error) {
	return Version, nil
}