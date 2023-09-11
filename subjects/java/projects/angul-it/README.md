## Angul-It

### Objectives

Develop a multiple-stage captcha web application using the latest version of Angular.

### Instructions

#### 1. Environment Setup

- Set up a new Angular application using the Angular CLI.
- Familiarize yourself with the generated project structure.
- Install any necessary dependencies to facilitate your development.

#### 2. Component Creation

- Establish vital components for the application, including but not limited to: `HomeComponent`, `CaptchaComponent`, and `ResultComponent`.

#### 3. Captcha Challenges

- Design the `CaptchaComponent` to challenge the user by identifying specific images from a grid.
- Empower users with the ability to revisit previous stages.

#### 4. Form Validation

- Implement robust form validation for every challenge. Ensure that users cannot transition to subsequent challenges without adequately completing the present one.

#### 5. State Management

- Harness Angular's state management to meticulously record user progression across challenges. Should the user opt to refresh the page, their progress must remain intact.

#### 6. Results Page

- Following the completion of all challenges, seamlessly redirect users to a results page. This page should present the challenge results and offer an option to embark on a new challenge.
- To maintain the challenge's integrity, prevent direct access to the results page without challenge completion. If a user attempts this, revert them to the challenge page.

### Bonus:

1. Diversify the challenge types and randomly allocate a unique set for each user session (e.g., image selection, math problem-solving, text input).
2. Incorporate animations to ensure fluid transitions between challenges.
3. Optimize the application for responsiveness, ensuring a seamless experience on both desktop and mobile platforms.
4. Construct comprehensive unit tests for your components.

### Testing

Your project will be extensively tested for the following aspects:

- Correctness of  components (`HomeComponent`, `CaptchaComponent`, and `ResultComponent`) implementation 
  
- form validation and captcha result.
  
- correctness of state management implementation and data saving in the browser.

