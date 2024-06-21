import { LitElement, html, css } from "lit";
import { property, customElement } from "lit/decorators.js";

@customElement("wowometer-component")
export class WowometerComponent extends LitElement {
  @property({ type: Boolean }) show = false;
  @property({ type: String }) studentID = "";
  @property({ type: Boolean }) wait = true;

  static styles = [
    css`
      :host {
        position: absolute;
      }
    `,
  ];

  tree = {
    Wait: {
      getHTML: () =>
        html`<section class="page center">
          <div class="header">
            <p class="title">Wait! 🖐</p>
            <p class="secondary">
              Please complete the Proof of Student Age & Identity first, as it
              may fulfill this requirement.
              <strong
                >If you've just submitted the Proof of Identity, please reload
                the page.</strong
              >
            </p>
          </div>
        </section>`,
    },
    ...UploadFlow,
  };

  render() {
    return html`
      <modal-component
        ?hide=${!this.show}
        .tree=${this.tree}
        .startingPage=${this.wait ? "Wait" : "UploadPage"}
        .inputParameters=${{
          uploader_documentName: `Custody / Guardianship Documentation`,
          uploader_maxFileSize: 8 * 1024 ** 2,
          uploader_supportedFiles: ".pdf",
          uploader_documentType: "documents",
          uploader_documentClass: "poc",
          uploader_student: this.studentID,
          uploader_steps: ["Uploading Document..."],
          uploader_subtitle:
            "There are many documents you can use as documentation. Please upload a supporting document.",
        }}
      >
      </modal-component>
    `;
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "wowometer-component": WowometerComponent;
  }
}
