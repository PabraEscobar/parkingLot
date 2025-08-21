Note: All the problems mentioned Below have been completed in the happy-friday branch [https://github.com/PabraEscobar/parkingLot/tree/happy-friday].
      Additionally, the solutions with interfaces and corresponding mocked interfaces for testing are available in the 
      mocking branch [https://github.com/PabraEscobar/parkingLot/tree/mocking]. Feel free to review them as needed.

      
Problem 1: Model a parking lot, Parking lots have a fixed capacity. Cars can be parked in a lot until it is full after which cars cannot be parked in the lot. 
           Unparking cars opens up space so cars can be parked again.
           
Problem 2: Check if car is parked, As a driver I would like to know if my vehicle is parked in the parking lot or not.

Problem 3: Notify Owner, As a parking lot owner I would like to be notified when the parking lot becomes full so that I can put up a sign saying the lot is full.
           As a parking lot owner I would like to be notified when the parking lot is available for parking again so that I can remove the sign.

Problem 4: Notify others, Security should to be notified when the parking lot becomes full so that he/she can redirect cars, 
           Traffic cop should be notified when the parking lot becomes full so that he/she can redirect cars.
           
Problem 5: Introducing Valet, An Attendant is responsible for parking and unparking cars from parking lots. When asked to park a car, 
           the attendant parks the car in the first lot with free space.

Problem 6: Valet / Attendant parks in lot with least number of cars When an attendant managing multiple parking lots is asked to park a car, 
           they should park in the lot with the least number of cars so that cars are evenly distributed between lots.
           Note: Attendants should still support parking in the first lot with free space. One should be able to choose between these two types of attendant.

Problem 7: Valet / Attendant parks in lot with most capacity , Extend the attendant to support parking in a lot with the most capacity. 
           Again, this is an extension and one should be able to choose between different types of parking styles.
