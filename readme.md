# GoKitchen
By Jeroenimoo0

GoKitchen is a project I made to learn GoLang. The project resembles a kitchen with different components all working together but running concurrently.

- **Supply**: A supply creates a certain resource every X seconds and adds this to the storage
- **Storage**: The storage stores all current existent resources. It however has a cap of 10 per resource and after this will block the supply.
- **Cooks**: Cooks create burgers using the resources from the storage. When there is a missing resource the cook will wait on it before he can continue finishing the burger. The cook also needs a customer before he can finish making his burger.
- **Customers**: Customers come in waiting for a burger. They however have a maximum time after which they will leave angry.

All these components depend on eachother, and will block work while waiting for another component to finish their job. This is all done using channels & goroutines.

I've created a simple frontend to showcase what is happening internally. Each node represents a component. It can have one of three colors
- **Green**: Currently doing nothing (Only customers will do this)
- **Orange**: Working on a job
- **Red**: Working on a job but being blocked by another component

A live version is running at https://gokitchen.haterd.net (Click on the '+' in the bottom row to trigger everything to work again)

## Misc
I've used Server Sent Events (SSE) to trigger live updates on the client.

## Credits
Also a little shoutout to [@YungTosti](https://github.com/YungTosti) for providing me with some basic HTML to start the frontend with.
