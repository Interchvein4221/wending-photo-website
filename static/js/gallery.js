const photos = window.galleryPhotos || [];

const lightbox = document.getElementById("lightbox");
const lightboxImage = document.getElementById("lightboxImage");
const photoCounter = document.getElementById("photoCounter");

const closeButton = document.getElementById("closeButton");
const prevButton = document.getElementById("prevButton");
const nextButton = document.getElementById("nextButton");
const scrollButton = document.getElementById("scrollButton");

let currentIndex = 0;
let touchStartX = 0;

function openPhoto(index) {
    if (!photos.length) return;

    currentIndex = index;
    lightboxImage.src = photos[currentIndex].src;
    lightboxImage.alt = photos[currentIndex].alt;
    photoCounter.textContent = `${currentIndex + 1} / ${photos.length}`;

    lightbox.classList.add("active");
    lightbox.setAttribute("aria-hidden", "false");
    document.body.style.overflow = "hidden";
}

function closePhoto() {
    lightbox.classList.remove("active");
    lightbox.setAttribute("aria-hidden", "true");
    document.body.style.overflow = "";
}

function showNext() {
    if (!photos.length) return;
    currentIndex = (currentIndex + 1) % photos.length;
    openPhoto(currentIndex);
}

function showPrevious() {
    if (!photos.length) return;
    currentIndex = (currentIndex - 1 + photos.length) % photos.length;
    openPhoto(currentIndex);
}

document.querySelectorAll(".photo-card").forEach((card, index) => {
    card.addEventListener("click", () => openPhoto(index));
});

closeButton.addEventListener("click", closePhoto);
nextButton.addEventListener("click", showNext);
prevButton.addEventListener("click", showPrevious);

lightbox.addEventListener("click", (event) => {
    if (event.target === lightbox) {
        closePhoto();
    }
});

document.addEventListener("keydown", (event) => {
    if (!lightbox.classList.contains("active")) return;

    if (event.key === "Escape") closePhoto();
    if (event.key === "ArrowRight") showNext();
    if (event.key === "ArrowLeft") showPrevious();
});

lightbox.addEventListener("touchstart", (event) => {
    touchStartX = event.changedTouches[0].screenX;
}, { passive: true });

lightbox.addEventListener("touchend", (event) => {
    const touchEndX = event.changedTouches[0].screenX;
    const distance = touchEndX - touchStartX;

    if (Math.abs(distance) < 50) return;

    if (distance < 0) {
        showNext();
    } else {
        showPrevious();
    }
});

scrollButton.addEventListener("click", () => {
    document.getElementById("gallery").scrollIntoView({
        behavior: "smooth"
    });
});
const musicButton = document.getElementById("musicButton");
const weddingMusic = document.getElementById("weddingMusic");

musicButton.addEventListener("click", () => {
    if (weddingMusic.paused) {
        weddingMusic.play();
        musicButton.textContent = "Ⅱ Пауза";
    } else {
        weddingMusic.pause();
        musicButton.textContent = "♫ Включить музыку";
    }
});

weddingMusic.addEventListener("ended", () => {
    musicButton.textContent = "♫ Включить музыку";
});
