# 🍽️ Dining Philosophers – Overview & Requirements

## 🏗 Project Overview

This project is a classic concurrency challenge implemented in Go: the Dining Philosophers problem. In this simulation, a group of philosophers sits around a circular table. Each philosopher alternates between thinking and eating. To eat, a philosopher must pick up the two forks (shared resources) adjacent to them. This setup creates potential issues such as deadlock and starvation if not handled correctly.

The project demonstrates key aspects of concurrent programming in Go, including:

- **Resource synchronization:** Managing access to shared forks using mutexes.
- **Deadlock prevention:** Implementing strategies (like an arbitrator or ordering of resource acquisition) to ensure that no philosopher gets stuck waiting indefinitely.
- **Goroutines & Channels:** Running each philosopher as a concurrent goroutine.
- **Fairness:** Ensuring that every philosopher eventually gets a chance to eat.

## 🎯 Challenge

Before diving into the solution, try implementing the Dining Philosophers problem yourself!

Your Task:
Create a concurrent simulation where:

✅ Each philosopher alternates between thinking and eating.
✅ A philosopher must acquire two forks (shared resources) to eat.
✅ Deadlock prevention strategies ensure that no philosopher gets stuck waiting indefinitely.
✅ Goroutines and mutexes are used to manage concurrency and resource sharing.

Bonus:
Implement different deadlock prevention strategies (e.g., arbitrator, ordered fork pickup).
Introduce metrics collection to track philosopher wait times and meals eaten.
Add a command-line flag to dynamically set the number of philosophers.
Visualize the simulation with a real-time graphical representation.
Once you've tried it, compare your approach with the provided solution to optimize further! 🚀

---

## **🔥 Features**

✅ **Concurrent Simulation:**

- Each philosopher runs in its own goroutine, simulating simultaneous thinking and eating.

✅ **Shared Resources:**

- Forks are modeled as mutex-protected resources shared between philosophers.

✅ **Deadlock Prevention:**

- Strategies are implemented to ensure that not all philosophers are waiting on each other (e.g., by imposing an order for picking up forks or using an arbitrator).

✅ **Randomized Behavior:**

- Random delays in thinking and eating simulate real-world unpredictability.

✅ **Logging:**

- Detailed logging helps trace each philosopher's actions (thinking, picking up forks, eating, and putting down forks).

✅ **Graceful Shutdown:**

- The simulation can be stopped gracefully after a fixed number of cycles.

---

## **🚀 Running the Project**

Ensure you have Go installed (`go version` to check). Then, run:

```sh
go run main.go
```

This command starts the simulation, where each philosopher alternates between thinking and eating, while the system prevents deadlock and ensures fairness.

## 🧪 Running Tests

To verify functionality and check for deadlocks or race conditions, run:

```sh
go test -race -v .
```

This command:

- Runs all unit tests
- Detects race conditions
- Provides verbose output for debugging

## 🔗 Related Concepts

- 📌 Goroutines & Channels:
  Foundations of concurrent programming in Go.

- 📌 Mutexes & Sync Mechanisms:
  Techniques for safely accessing shared resources.

- 📌 Deadlock & Starvation Prevention:
  Strategies to ensure all concurrent processes make progress.

- 📌 Resource Scheduling & Fairness:
  Ensuring that shared resources are used fairly among competing goroutines.

- 📌 Simulation & Randomization:
  Using randomized delays to mimic real-life concurrency scenarios.

## 📸 Future Enhancements

- 📌 Dynamic Philosopher Count:
  Allow the number of philosophers and forks to be configurable via command-line flags.

- 📌 Multiple Deadlock Prevention Strategies:
  Implement and compare different strategies (e.g., arbitrator-based vs. ordered fork pickup).

- 📌 Metrics Collection:
  Track how many times each philosopher eats and how long they wait.

- 📌 Graphical Visualization:
  Visualize the simulation progress and resource usage in real-time.

- 📌 Graceful Termination:
  Add a mechanism to stop the simulation cleanly after a certain time or number of cycles.
- 📌 Distributed Dining Philosophers:
  Extend the simulation to a distributed setting with multiple tables and shared resources.
- 📌 Add more test

## 💡 Author & Contributions

🛠 Created by Oluyemi Olususi
📢 Open to contributions & discussions! Feel free to fork, improve, and share feedback.

✅ What’s Next?
Now that this project is complete, I'll be working on additional concurrency challenges in Go. Stay tuned for:
