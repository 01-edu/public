## Secure Messaging App

Messaging apps have become an essential part of our daily lives, allowing us to easily communicate with friends, family, and colleagues. Over the years, they evolved significantly, from simple text messaging to video calling and more. However, as messaging apps have become more popular, concerns about privacy and security have also increased. Many messaging apps have been criticized for their lack of security and the potential for users private data to be compromised.

### Instructions

Build a secure mobile messaging app using `Flutter` and `Firebase` (or similar technologies). It should support both iOS and Android. The app should have **biometric authentication**, **end-to-end encryption**, and other security features to protect user data. Users should be able to send text messages, images and videos.

**Make sure you implement the following features:**

- **Login/Signup**: Implement a login/signup page and the necessary functionality for user authentication and account creation.
- **User profile**: Users must be able to create a profile with a profile picture, username, and other information.
- **Search users**: Users must be able to search for other users by username or other information.
- **Add contacts**: Users must be able to add contacts by searching for their username or scanning a QR code.
- **Messaging**: Users must be able to send text messages, images, and videos. Apart from that you must implement these features:
  - **read receipts**
  - **typing indicators**
  - **edit** or **delete** messages.
- **Secret chats**: Users must be able to start encrypted one-on-one chats with the same functionality as regular chats with one exception: all messages have to be encrypted using an end-to-end algorithm.

### Bonus:

- **Group messaging**, where users are able to create or join group chats with multiple users having at least the same functionality as regular chats.

- **Push notifications**, where users can receive notifications of new messages or friend requests using the firebase_messaging package.

#### Suggested Packages for Flutter and Firebase:

- firebase_auth
- cloud_firestore
- firebase_storage
- cached_network_image
- image_picker
- video_player
- flutter_secure_storage
- encrypt
- local_auth
- firebase_messaging
