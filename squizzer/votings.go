package squizzer
import(
	"fmt"
	"squizz/tools"
)

type VotingOption struct{
	id string
	name string
	voteCount int
}

type Voting struct{
	id string
	name string
	choices []VotingOption
	winner VotingOption
	status string
}

var VotingStatus = []string{"pending","open","closed","archived"}

func ValidateVoting(voting Voting) error{
	
	if (tools.ArrayFind(VotingStatus,voting.status)<0){
		return fmt.Errorf("Invalid status")
	}

	return nil
}

func GetAllVotings(){
	// var allVotings[]
}