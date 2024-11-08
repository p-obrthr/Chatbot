let messagesContainer = document.getElementById("messagesContainer");
let userInput = document.getElementById("userInput");
let isLoading = false;

const messages = [
    // first partner message init
    { content: "Hello, I'm your chatbot. May I can help you?", isUser: false }
];

function renderMessages() {
    messagesContainer.innerHTML = "";

    // Render each message
    messages.forEach(message => {
        const messageDiv = document.createElement("div");
        messageDiv.classList.add("message", message.isUser ? "user-message" : "partner-message");
        messageDiv.textContent = message.content;
        messagesContainer.appendChild(messageDiv);
    });

    if (isLoading) {
        const loadingAnimation = document.createElement("dotlottie-player");
        loadingAnimation.src = "https://lottie.host/8f2bf6c2-5e60-49df-8bed-6149130f98df/UnGJDGRQ4r.lottie";
        loadingAnimation.className = "lottie-animation";
        loadingAnimation.style.width = "80px";
        loadingAnimation.style.height = "80px";
        loadingAnimation.setAttribute("background", "transparent");
        loadingAnimation.setAttribute("speed", "1");
        loadingAnimation.setAttribute("loop", "");
        loadingAnimation.setAttribute("autoplay", "");
        messagesContainer.appendChild(loadingAnimation);
    }
    setTimeout(scrollToBottom, 100);
}

function sendMessage() {
    const text = userInput.value.trim();
    if (text === "") return;

    messages.push({ content: text, isUser: true });
    userInput.value = "";
    isLoading = true;
    renderMessages();

    // dummy  answer
    setTimeout(() => {
        isLoading = false;
        messages.push({ content: "Hello", isUser: false });
        renderMessages();
    }, 3000);
}

function scrollToBottom() {
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
}

function handleEnter(event) {
    if (event.key !== "Enter") return
    sendMessage();
}

document.getElementById("userInput").addEventListener("keydown", function(e) {
    if (e.key === "Enter") sendMessage();
});

// init
renderMessages();
