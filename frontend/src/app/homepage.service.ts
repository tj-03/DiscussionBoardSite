import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class HomepageService {

  getAllPosts() {
    return this.http.get<{id: string, author_id: string, content: string}[]>('http://localhost:8080/api/posts');
  }
}
