import { Component, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-sidebar',
  standalone: true,
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
  imports: [CommonModule],
})
export class SidebarComponent {
  active = 'home';
  @Output() tabChange = new EventEmitter<string>();

  folders = [
    { name: 'Home', icon: 'home', tab: 'home' },
    { name: 'AI Assistant', icon: 'chat', tab: 'assistant' },
    { name: 'About', icon: 'info', tab: 'about' }
  ];

  changeTab(tab: string) {
    this.active = tab;
    this.tabChange.emit(tab);
  }
}
