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
        messageDiv.innerHTML = markdownToHtml(message.content); 
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

function markdownToHtml(text) {
    text = text.replace(/(\*\*|__)(.*?)\1/g, '<strong>$2</strong>'); 
    text = text.replace(/(\*|_)(.*?)\1/g, '<em>$2</em>'); 
    text = text.replace(/^\#\s(.*)/gm, '<h1>$1</h1>'); 
    text = text.replace(/^\#\#\s(.*)/gm, '<h2>$1</h2>'); 
    text = text.replace(/^\#\#\#\s(.*)/gm, '<h3>$1</h3>'); 
    text = text.replace(/^\- (.*)/gm, '<ul><li>$1</li></ul>');

    return text;
}


async function sendMessage() {
    const text = userInput.value.trim();
    if (text === "") return;

    messages.push({ content: text, isUser: true });
    userInput.value = "";
    isLoading = true;
    renderMessages();

    try {
        const response = await fetch("http://localhost:4000/api/messages", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ question: text })
        });

        if (response.ok) {
            const data = await response.json();
            isLoading = false;
            console.log(data)
            messages.push({ content: data.answer, isUser: false });
            renderMessages();
        } else {
            isLoading = false;
            messages.push({ content: "error: response problem", isUser: false });
            renderMessages();
        }
    } catch (error) {
        isLoading = false;
        messages.push({ content: "error: server problem", isUser: false });
        renderMessages();
    }
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
