import { Component, Input, OnChanges } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { ChatService } from './chat.service';
import { ChangeDetectorRef } from '@angular/core';

@Component({
  selector: 'app-chat',
  standalone: true,
  imports: [CommonModule, FormsModule], // only keep Angular basics
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.css']
})
export class ChatComponent implements OnChanges {
  @Input() defaultPrompt: string = '';
  messages: { sender: string; text: string; sources?: any[] }[] = [
    { sender: 'PAIGE', text: 'Welcome! How can I help with your concern?' }
  ];
  userInput: string = '';

  constructor(
    private chatService: ChatService,
    private cdr: ChangeDetectorRef
  ) {}

  ngOnChanges() {
    this.userInput = this.defaultPrompt; // pre-fill input box
    this.sendMessage(); // automatically send the default prompt as a message
  }

  sendMessage() {
    if (!this.userInput.trim()) return;

    const question = this.userInput;

    this.messages.push({
      sender: 'You',
      text: question
    });

    this.userInput = '';

    this.chatService.askQuestion(question).subscribe({
      next: (res) => {
        console.log(res);

        this.messages.push({
          sender: 'PAIGE',
          text: res.answer,
          sources: res.sources
        });

        console.log(this.messages);
        this.cdr.detectChanges();
      },

      error: (err) => {
        this.messages.push({
          sender: 'PAIGE',
          text: 'Error: ' + err.message
        });
        this.cdr.detectChanges();
      }
    });
  }
}
