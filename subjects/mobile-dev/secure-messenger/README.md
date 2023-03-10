# Secure Messaging App

### Introduction

Messaging apps have become an essential part of our daily lives, allowing us to easily communicate with friends, family, and colleagues. Over the years, they evolved significantly, from simple text messaging to video calling and more. However, as messaging apps have become more popular, concerns about privacy and security have also increased. Many messaging apps have been criticized for their lack of security and the potential for users' private data to be compromised.


### Objective

Build a secure mobile messaging app using Flutter and Firebase (or similar technologies). It should support either iOS, Android or both. The app should have biometric authentication, end-to-end encryption, and other security features to protect user data. Users should be able to send text messages, images, and videos, and create or join group chats.


### Instructions for Flutter and Firebase as a reference:

Set up a new Flutter project and integrate Firebase authentication and Cloud Firestore as a backend. Use the Firebase Console to create a new Firebase project, add appropriate configuration files.

To enable user authentication implement Firebase Authentication and secure sign-in. Users should be able to sign up, log in, and log out of the app using email and password authentication. Additionally, implement biometric authentication using the local_auth package to allow users to log in using their fingerprint or face ID.

Implement end-to-end encryption for secret chats using the encrypt or webcrypto packages. Use a symmetric encryption algorithm like RSA, and generate a new public and private key pair for each pair of users for each secret chat. The public keys should be securely shared between the sender and receiver.



### Implement the following features:

- User profiles: Users should be able to create a profile with a profile picture, username, and other information.
- Search users: Users should be able to search for other users by username or other information.
- Add contacts: Users should be able to add contacts by searching for their username or scanning a QR code.
- Messaging: Users should be able to send text messages, images, and videos, with features such as read receipts, typing indicators, and the ability to edit or delete messages.
- Secret chats: Users should be able to start encrypted one-on-one chats with the same functionality as regular chats with one exception: all messages have to be encrypted using an end-to-end algorithm.



### Bonus:

Group messaging: Users should be able to create or join group chats with multiple users having at least the same functionality as regular chats.

Push notifications: Users should receive notifications of new messages or friend requests using the firebase_messaging package.

## Suggested Packages for Flutter and Firebase:

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