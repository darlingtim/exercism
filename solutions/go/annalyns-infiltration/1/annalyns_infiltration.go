package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
    /* assign what is assigned the value of KnightIsAwake function parameter to the knightIsAwake variable. 
    no need to declare it as it has been declared at the function declaration level.*/
    knightIsAwake = knightIsAwake

    /*if value of knightisAwake is true, then return false. else he is sleeping so we return true because he can be attacked.*/
    if knightIsAwake{
        return false
    }else{
        return true
    }
	panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
    // one of the threqe must be awake
    if knightIsAwake == true || archerIsAwake == true || prisonerIsAwake == true {
        return true
    }else {
        return false
    }
	panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if archerIsAwake == false && prisonerIsAwake == true {
        return true
    }else{
        return false
    }
	panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    //only the prisoner must be awake or the dog should be present while the archer is sleeping
    if (prisonerIsAwake == true && knightIsAwake == false &&  archerIsAwake == false) || (petDogIsPresent == true && archerIsAwake == false) {
        return true
    }else{
        return false
    }
	panic("Please implement the CanFreePrisoner() function")
}
