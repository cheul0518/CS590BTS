/*



*/


package main



import (

	"fmt"
	"strconv"
	"string"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"

)

var logger = shim.NewLogger("example_cc0")

// Car Component
type CarComponent struct {
    retired     bool
    owner       string      // entity: "ROLE_TYPE.ROLE_NAME"
	carID	string

}

func (t * CarComponent) Init(stub shim.ChaincodeStubInterface) peer.Response {
    // No action, because there is no components at the very beginning
	return nil
}


// Invoking the correct function
func (t * CarComponent) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    
    fn, args := stub.GetFunctionAndParameters()

    var result string
    var err error

    if fn == "AddComponent" {

		result, err = t.AddComponent(stub, args)

	} else if fn == "TransferComponent" {

		result, err = t.TransferComponent(stub, args)
            
	} else if fn == "MountComponent" {

		result, err = t.MountComponent(stub, args)

	} else if fn == "ReplaceComponent" {

		result, err = t.ReplaceComponent(stub, args)

	} else if fn == "RecallComponent" {

		result, err = t.RecallComponent(stub, args)
    
	} else if fn == "CheckComponent" {

		result, err = t.CheckComponent(stub, args)
                        
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(result))
        
}



// AddComponent(Role, ComponentID)
func (t *CarComponent) addComponent(stub shim.ChaincodeStubInterface, args []string) peer.Response {


	var A string // ComponentID
	var err error

	if len(args) != 2 {
		return shim.Errorf("Incorrect number of arguments. Expecting two")
	}

	if args[0] != "Supplier" {
		return shim.Errorf("Incorrect role. It's supposed to be Supplier")
	}


	if len(args[1]) != 9 {
		return shim.Errorf("Incorrect componentID length")
	}


	A = args[1]

	component, err := stub.GetState(A)

	if err != nil {
		return shim.Errorf(err.Error())
	} else if component != nil && component.retired {
		return shim.Errorf("Error: Using retired component")
	}

	component = CarComponent{false, "Supplier", ""}

	err = stub.PutState(A, []byte(component))

	return shim.Success(nil)	// need to check what to return
	

}


// TransferComponent(Role, New Owner, ComponentID)
func (t *CarComponent) transferComponent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	
	var A,B string // new owner, componentID
	
	if len(args) != 3{
		return shim.Errorf("Incorrect number of arguments. Expecting three")
	}
	
	if args[0] == "Cars" {
		return shim.Errorf("Incorrect role. It's supposed to be Supplier/Manufacturer/Authorized dealer")		
	}

	if len(args[2]) != 9 {
		return shim.Errorf("Incorrect componentID length")
	}

	A = args[1]
	B = args[2]

	// delete component ID in ledger
	err  := stub.DelState(B)
	if err != nil {
		return shim.Errorf(err.Error())
	}

	// add componenet ID with new owner in ledger
	component := CarComponent{false, args[0], ""}
	err = stub.PutState(B, []byte(component))
	if err != nil {
		return shim.Errorf(err.Error())
	}	

	return shim.Success(nil)
}

// MountComponent(Role, ComponentID, CarID)
func (t *CarComponent) mountComponent(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var A,B string

	if len(args) != 3{
		return shim.Errorf("Incorrect number of arguments. Expecting three")
	}
	if len(args[1]) != 9 {
		return shim.Errorf("Incorrect componentID length")
	}
	if args[0] != "Manufacturer" {
		return shim.Errorf("Incorrect role. It's supposed to be Manufacturer")
	}

	A = args[1]
	B = args[2]

	// delete component ID in ledger
	err  := stub.DelState(A)
	if err != nil {
		return shim.Errorf(err.Error())
	}

	// add componenet ID with new owner in ledger
	component := CarComponent{false, args[0], args[2]}
	err = stub.PutState(A, []byte(component))
	if err != nil {
		return shim.Errorf(err.Error())
	}	

	return shim.Success(nil)

	
}

func (t * CarComponent) CheckComponent(stub shim.ChaincodeStubInterface, args []string) (bool, error) {

    if len(args) != 2 {
        return false, shim.Errorf("Incorrect argument: need the role and ComponentID")
    }

    entity := args[0]

    if !strings.EqualFolds(entity, "Car") {
        return false, shim.Errorf("%s cannot access this function: only 'Car' allowed", entity)
    }

    componentID := args[1]

    if len(componentID) != 9 {
        return shim.Errorf("Error: only accept nine alphanumberic componentID")
    }

    component := stub.GetState(componentID)

    return !component.retired       // return true it not retired
                                    // return false if retired
}

func (t * CarComponent) RecallComponent(stub shim.ChaincodeStubInterface, args []string) (error) {
    
    if len(args) != 2 {
        return shim.Errorf("Incorrect argument: need the role and ComponentID")
    }

    entity := args[0]

    if !strings.EqualFolds(entity, "Manufacture") {
        return shim.Errorf("%s cannot access this function: only 'Manufacture' allowed", entity)
    }

    componentID := args[1]
    component := stub.GetState(componentID)
    retired := component.retired

    if len(componentID) != 9 {
        return shim.Errorf("Error: only accept9 digit componentID")
    }

    if retired {
        return shim.Errorf("Cannot recall retired component")
    }

    // update the owner to be the Manufacture
    err := peer.PutState(componentID, CarComponent{false, "Manufacture", ""})

    return err

}


func (t * CarComponent) ReplaceComponent(stub shim.ChaincodeStubInterface, args []string) (error) {
    
    if len(args) != 3 {
        return shim.Errorf("Incorrect argument: need the role, ComponentID and CarID")
    }

    entity := args[0]

    if !strings.EqualFolds(entity, "Manufacture") &&
       !strings.EqualFolds(entity, "Dealer") {
        return shim.Errorf("%s cannot access this function: only 'Manufacture' and 'Dealer' allowed", entity)
    }

    componentID := args[1]
    carID := args[2]

    if len(componentID) != 9 {
        return shim.Errorf("Error: only accept 9 digit componentID")
    }

    component := stub.GetState(componentID)
    retired := component.retired

    if retired {
        return shim.Errorf("Error: component retired")
    }

    // remove the component
    err := stub.DelState(componentID)

    return err

}


func main() {
	err := shim.Start(new(CarComponent))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}
