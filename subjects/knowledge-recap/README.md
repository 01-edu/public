## Knowledge Recap

### Overview

From this point on, you’re expected to actively reflect on your learning journey after every project you complete. The **Knowledge Recap** is your personal documentation space, a place where you’ll track what you’ve learned, how you overcame challenges, and what skills you’ve developed.

You’ll start by creating a dedicated Gitea repository, where your very first `README.md` will introduce your personal goals for the program. Then, after each project, you’ll add a short reflection file answering key questions about your experience.

This project will help you become more self-aware, build a strong technical narrative, and make your progress visible to others (peers, mentors, future employers). It will also be taken into consideration during a final face-to-face exam with the jury at the end of this program, opening many opportunities to you.

---

### Role Play

You are a developer building your portfolio and sharpening your skills. You’ve just finished a project, and now your lead asks you:

> "What did you learn? What went wrong? How did you fix it? What are you better at now?"

The Knowledge Recap is a place where you can answer these questions clearly, similar to how you would in a professional development role where formal documentation is required.

---

### **Instructions**

### 1. Create Your Repository

Create a new repository on Gitea and name it:

`knowledge-recap`

### 2. Write Your Introduction

In the `README.md` file, write about yourself:

- Who are you as a learner?
- Why did you join this program?
- What are your goals?
- What do you hope to achieve by the end?

This file will serve as your **anchor,** the starting point of your journey.

### 3. After Every Project

For **every** future project, create a new markdown file inside your `knowledge-recap` repository.

Use the project’s name as the file name (for example: `networking-shell.md` or `web-scraper.md`), and answer the following questions:

1. What did you learn?
2. How did you learn it? (resources, people, trial-and-error…)
3. What issues or blockers did you face?
4. How did you solve or work around them?
5. What skills did you improve?
6. If you had to redo the project, what would you change?
7. How does this project relate to your long-term goals?

> The submission should include relevant screenshots from the submission / project in question.

---

### Project Repository Structure

```console
/knowledge-recap
├── README.md              # Your personal goals and intro
├── deep-in-net.md        # Recap for the Shell Basics project
├── .... throughout the program, the rest of the Readme files will go here
├── **project-name.md**
├── /screenshots
│   ├── shell-basics.png   # Screenshot(s) of the shell project output
│   ├── web-scraper.png    # Screenshot(s) showing working scraper
│   └── ...

```

**README.md:** The starting point: who you are, what you want, and why you’re here.

**project-name.md:** A reflection after each project. Each file answers the 7 questions above, and you will add it after each project you finish from now on.

---

### Tips

- **Be specific**. Don’t say “I learned new things.” Say, “I learned to use `grep` to filter logs.”
- **Be real**. Don’t be afraid to talk about mistakes or confusion. Growth starts with honesty.
- **Commit regularly**. Use meaningful commit messages like `docs: added recap for shell basics`.
- **Make it yours**. Feel free to include screenshots, terminal outputs, or even links to cool resources you found.
- **Make it professional**. Use correct grammar and proofread your writing.

---

> NOTE: This project never really ends. It's your way to track your evolution as a developer. By the end of the program, this repository will be proof of how far you’ve come, not just in code, but in thinking, learning, and building and eventually may change your future.
