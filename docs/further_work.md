# Further Work

## Expand the service to use an eventstore
Currently our service is stateless which makes the use of an eventstore a bit challenging. 
Without extending the functionality we could change the interactions between the client and the server to happen through the use of **Commands** and **Events**. The client could issue a command **AddNumbers** and the service could create the event **NumbersAdded** once it is done.

If we had application state we could use an eventstore that our service would consume to recreate its state from a sequence of events, this is Event Sourcing. 
An example would be if we wanted to store a total, then we could use events **TotalIncreased**, **TotalReduced**. Storing all the events in a store would mean that even if our service lost its state it could recreate it by replaying all the events. 