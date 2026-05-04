import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ChatService {
  // private apiUrl = 'https://paige-be.onrender.com/ask';
  private apiUrl = 'http://localhost:8080/ask';

  constructor(private http: HttpClient) {}

  askQuestion(question: string): Observable<{answer: string, sources: any[]}> {
    const body = { question };
    return this.http.post<{answer: string, sources: any[]}>(this.apiUrl, body);
  }
}
