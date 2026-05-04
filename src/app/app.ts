import { Component } from '@angular/core';
import { ChatComponent } from './chat/chat.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-root',
  templateUrl: './app.html',
  styleUrls: ['./app.scss'],
  imports: [CommonModule, ChatComponent, SidebarComponent, AboutComponent, HomeComponent]
})
export class App {
  activeTab: string = 'home';
  chatPrompt: string = '';

  setActiveTab(tab: string) {
    this.activeTab = tab;
  }

  handleOpenChat(prompt: string) {
    this.activeTab = 'assistant';
    this.chatPrompt = prompt;
  }

  contactDev() {
    window.location.href = 'https://docs.google.com/forms/d/e/1FAIpQLSf5I2-y2HjxptBO3SZWFgBmctWmCGYoecVPslKdDmUI5SlWDg/viewform';
  }
}
