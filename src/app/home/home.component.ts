import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-home',
  standalone: true,
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
  imports: [CommonModule]
})
export class HomeComponent {
  // EventEmitter that sends a prompt string to parent
  @Output() openChat = new EventEmitter<string>();

  popupDoc: string | null = null;
  safeUrl: SafeResourceUrl | null = null;

  constructor(private sanitizer: DomSanitizer, private router: Router) {}

  openDoc(doc: string) {
    this.popupDoc = doc;

    // Option A: embed Google Doc via iframe
    const url = this.getDocUrl(doc);
    if (url) {
      this.safeUrl = this.sanitizer.bypassSecurityTrustResourceUrl(url);
    }
  }

  closePopup() {
    this.popupDoc = null;
  }

  // Map document names to Google Doc embed URLs
  getDocUrl(doc: string) {
    const urls: Record<string, string> = {
      'Employment Guidelines': 'https://docs.google.com/document/d/1lZZD3CuLhl6Rby1LbgaXSzvp4EXPzgCeKgSjcFBlywY/preview',
      'How to File a Concern': 'https://docs.google.com/document/d/1bZft5Qfhnd_NHLS5lB8UdmcVx0Km4r6L8A7pbylNeYU/preview',
      'Policies About Hiring': 'https://docs.google.com/document/d/15N0TZ-imSAuuV-rE9_5tR6CyKnzw-oOD0bpidtAvqS4/preview',
      'Capabilities Upon Employment': 'https://docs.google.com/document/d/1lbUKlGyQJEYOiwEI9SXQD4r35FsCwqM-rkY_hKBM-74/preview',
      'How to Raise a Ticket': 'https://docs.google.com/document/d/1elecTVyTw9M1JrOHvdM9LIkSHIm3AgvkFLa49AoUO_c/preview'
    };
    return urls[doc];
  }

  goToChat(docName: string) {
    const defaultPrompt = `Give me a details about ${docName}`;
    this.openChat.emit(defaultPrompt);
  }
}
