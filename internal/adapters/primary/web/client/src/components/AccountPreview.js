export default class AccountPreview extends HTMLElement {
    constructor() {
        super();
    }

    connectedCallback() {
        const name = this.querySelector(".account-name");
        const text = name.querySelector("h5");
        text.classList.toggle("group-hover:animate-marquee", name.clientWidth < text.clientWidth);

        console.log(text.innerText, name.clientWidth, text.clientWidth, name.clientWidth < text.clientWidth);
    }
}

