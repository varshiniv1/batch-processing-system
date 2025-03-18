# Batch Processing System

This project is a **Distributed Batch Processing System** built using **Golang**, **Google Cloud Pub/Sub**, and **Google Cloud Platform (GCP)**. It allows users to submit batch jobs via an API, processes them using a worker, and logs the results.

---

## Features
- **Job Submission API**: Submit jobs via a RESTful API.
- **Job Queue**: Jobs are queued using Google Cloud Pub/Sub.
- **Worker**: Processes jobs and logs the results.
- **Scalable**: Designed to handle large-scale batch processing.

---

## How to Run the Project

### Prerequisites
1. **Google Cloud Account**: Set up a GCP project and enable the Pub/Sub API.
2. **Google Cloud SDK**: Install and configure the SDK.
3. **Golang**: Install Go (version 1.20 or higher).

### Steps to Run

#### 1. Clone the Repository
```bash
git clone https://github.com/varshiniv1/batch-processing-system.git
cd batch-processing-system
