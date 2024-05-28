# Cross Site Request Forgery and its prevention in GoLang

In this lab we will be exploring Cross-Site Request Forgery attacks, or CSRF for short, in the context of GoLang web applications.
For our implementation, we have a simple web application that serves as a login page.

## Instructions

### 1. Let us spin up our vulnerable web application

```bash
docker build -t csrf-app .
```

> This should take about a minute, in the mean time let us talk about the how this exploit works.
> Ideally, only the developer-defined endpoint must act as a sign-in page.
> However, in case of CSRF attacks, this request is forged by a malicious webpage that serves to simulate our original end-point.
> In most cases, these links can be disguised as spam or ads that POST to the end-point with hidden fields and on-click submit options leaving the victim entirely unaware.
> These attacks can result in loss of confidential data, theft of resources, user impersonation, etc.

```bash
docker run --name csrf-cont -p 9090:9090 -p 9091:9091 -d csrf-app
```

> This command runs our csrf-app image on the container named csrf-cont while exposing the ports 9090-9091 and binding the same from the container to the host.

### 2. Connecting to the container and accessing the application

> Once the container is running, we use `serverip` to connect to it:

```bash
serverip
```

> This should give us the public IP Address of the server that we need to connect to.
> Point your browser to `<serverip>:9090/` to access a simple login page.

### 3. Using the application

> The application exposes a login page that takes a username and password.
> These are not being stored anywhere but only reflected back from the login page.
> Once you've entered these details and posted them, you will find that the page reloads and reflects at the top of the page.
> To showcase that this application is vulnerable to CSRF attacks, let us start another server that will spoof the identity of our original webpage.

### 4. Showcasing CSRF Attacks

> To spoof this webpage, let us start another server and serve a similar page through it.
> First, let us start a bash shell within the container:

```bash
docker exec -it csrf-cont bash
```

> Once you have access to a shell within the container, you can navigate to the ./hackerprog:

```bash
ls
cd hackerprog
ls
```

> Within this directory, you should see an application by the name hacker.go in this space.
> To run this file, do `go run hacker.go`
> This starts a server that listens on the PORT 9091
> Point your browser towards this port by going to `<serverip>:9091/`
> You will be met with a webpage congratulating you for winning the jackpot and asking you to fill in your details.
> Once you fill in your credentials and POST them, you will be redirected to the original webpage that now reflects the credentials that you posted to the fake website.

### 5. Defense against CSRF attacks

> The consequences of what you've just seen are huge. If a malicious actor is able to spoof your original webpage with hidden fields that contain information like payment details, they could steal information or even actual money.
> To defend against these attacks, we use CSRF tokens that are randomly generated secure strings. These tokens can be embedded into forms, HTML headers, etc.

Let us run a secure variant of our application to defend against these attacks:
> Start by replacing the contents of `main.go` with the `safe.txt` and save the file.
> Build the docker image with the new secure variant of the application.

```bash
docker build -t csrf-app .
```

> While that is happening, a few things to draw your attention to in the new application:
> Echo framework has built-in functionality that lets us use CSRF tokens while handling POST requests.
> These are embedded into the context and retrieved by the request handler function.
> In case the CSRF token is not verified, the request is turned away with a HTML 4XX error.

Start the container:

```bash
docker run --name csrf-cont -p 9090:9090 -p 9091:9091 -d csrf-app
```

> Just like before, point your browser to `<serverip>:9090/` and check if the application is working as before.
> To test if it is vulnerable to CSRF Attacks, we run another the hacker application as before.
> Execute a bash shell within the same container by using these commands:

```bash
docker exec -it csrf-cont bash
```

```bash
cd hackerprog
go run hacker.go
```

> This will start another server that listens on the PORT 9091. Point your browser towards `<serverip>:9091/`
> Try posting credentials from the fake website and note the results.

## Teardown

```bash
docker stop csrf-cont
```
