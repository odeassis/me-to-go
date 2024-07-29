package main

func nextPrime(number int) (result int, err error) {

	if number < 2 {
		return 2, nil
	}

	for i := 2; i <= number; i++ {
		if number%i == 0 && i != number {
			return nextPrime(number + 1)
		} else {
			if number == i {
				return number, err
			}
		}
	}

	return result, err
}

func main() {

}
