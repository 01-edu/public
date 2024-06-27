## kaquiz

### Instructions

The mobile app you will need to develop is designed to allow users to connect with their friends and track their last known location. Upon logging into the app via email, users need to be able to search for their friends by email and send friend requests. Once two users have become friends within the app, they will be able to see each other's last known location when the app is turned on.

The app needs to use location tracking technology to determine the last known location of each user. This technology will be able to determine the users location within a certain degree of accuracy, depending on the capabilities of the device being used.

Overall, the app is designed to be a fun and interactive way for users to stay connected with their friends and see where they've been. Whether you're meeting up for a night out on the town or just keeping tabs on your friends whereabouts, this app is sure to be a hit with anyone who loves mobile technology.

This app is fullstack, therefore you will have to implement both backend and frontend for the project. For the frontend part, make sure to add following functionality:

- Authentication via email:

  - Implement a secure and user-friendly email authentication mechanism.

- Ability to search for friends via email:

  - Design a search feature that enables users to find their friends efficiently.

- Sending friends request:

  - Create a user-friendly interface to send friend requests and implement the appropriate validation to avoid duplicate requests or invalid email addresses.

- Denying/Accepting friend requests:

  - Provide clear and intuitive options for users to accept or decline incoming friend requests.

- Deleting friends from the list:

  - Design a straightforward process for users to remove friends from their friend list.

- Viewing friends last known location on the map:

  - Develop a visually appealing and informative display to show your friends locations incorporating `maps` or `geolocation` services to provide accurate location data.

To keep users location up to date, each time the app is open, current location must be sent to the server every 5 seconds.

Backend implementation and Swagger details:

- Refer to the provided [swagger.yml](swagger.yml) file for backend implementation details.
- Ensure the backend functionality aligns with the frontend requirements.

#### Hints

To ensure accurate location tracking and smooth backend integration, keep the following points in mind:

- Location tracking technology:

  - Utilize appropriate location tracking technologies compatible with various devices.
  - Account for different levels of accuracy based on the capabilities of the user's device.

- Sending location updates to the server:

  - Implement a mechanism that sends the user's current location to the server.
  - Set the interval to every 5 seconds when the app is open to maintain up-to-date location data.
