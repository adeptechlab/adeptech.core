package enumerations

type Object string

var (
	ObjectTypeFeedback Object = "feedback"
	ObjectTypeAccount  Object = "account"
	ObjectTypeUser     Object = "user"
)

func (t Object) String() string {
	return string(t)
}
