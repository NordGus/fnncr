export default class DashboardPreview extends HTMLElement {
    constructor() {
        super();
    }

    connectedCallback() {
        const name = this.querySelector(".name");
        const text = name.querySelector("h5");
        text.classList.toggle("group-hover:animate-marquee", name.clientWidth < text.clientWidth);
    }
}

