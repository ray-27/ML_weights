# ML Weights

This project is a SaaS application designed to track and visualize machine learning experiments, similar to Weights & Biases. It provides real-time logging of training and validation metrics, inference result storage, and a user-friendly web dashboard to visualize model performance. The platform is built using **Golang** with the **Gin framework**, deployed on **AWS EC2**, and integrates **GitHub OAuth** for secure user authentication.

## 🚀 Features
- Real-time tracking of training and validation metrics using **WebSocket architecture**
- Secure authentication via **GitHub OAuth**
- Inference result storage and visualization
- User-friendly web dashboard for ML performance insights
- Scalable backend using **Golang** with **Gin framework**
- Efficient request handling and compute operations with **concurrency techniques**

## 🛠️ Tech Stack
- **Golang**
- **Gin framework**
- **AWS EC2**
- **DynamoDB** (for data storage)
- **Memcache** (for caching)
- **Route 53** (for domain management)
- **Lambda functions** (for serverless tasks)
- **WebSockets** (for real-time updates)
- **GitHub OAuth** (for authentication)

## 🔧 Setup and Installation
### Prerequisites:
- Go (1.18+)
- AWS account
- GitHub account for OAuth setup

### Steps:
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ml-tracking-saas.git
   cd ml-tracking-saas
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up environment variables:
   - `GITHUB_CLIENT_ID` - GitHub OAuth client ID
   - `GITHUB_CLIENT_SECRET` - GitHub OAuth client secret
   - `AWS_ACCESS_KEY` - AWS access key
   - `AWS_SECRET_KEY` - AWS secret key

4. Run the server:
   ```bash
   go run main.go
   ```

## 📡 Deployment
The platform is deployed on **AWS EC2**. Follow these steps to deploy:
1. Launch an EC2 instance and SSH into it.
2. Install Go on the instance.
3. Clone the repository and set up environment variables.
4. Run the server using:
   ```bash
   go run main.go
   ```

## 📊 Usage
- Log into the platform using your **GitHub account**.
- Track your ML experiments by logging training/validation metrics via API endpoints.
- Visualize metrics and inference results in real-time on the web dashboard.

## ⚙️ API Endpoints
| Method | Endpoint            | Description                  |
|--------|---------------------|------------------------------|
| POST   | `/login`            | GitHub OAuth login           |
| POST   | `/log-metrics`      | Log training/validation data |
| GET    | `/get-results`      | Retrieve inference results   |
| POST   | `/store-results`    | Store inference results      |

## 📚 Concurrency Techniques
To ensure seamless handling of requests and compute operations, the platform utilizes:
- **Goroutines** for concurrent task execution
- **Channels** for safe communication between Goroutines
- **Mutexes** to handle shared resources and avoid race conditions

## ☁️ AWS Services Used
- **EC2**: For hosting the backend server
- **DynamoDB**: For storing metrics and results
- **Memcache**: For caching frequently accessed data
- **Route 53**: For managing the domain
- **Lambda**: For serverless operations and tasks

## 🤝 Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## 📄 License
This project is licensed under the MIT License.

## 🙋‍♂️ Contact
For any queries, feel free to reach out:
- **Email**: your.email@example.com
- **GitHub**: [yourusername](https://github.com/yourusername)

