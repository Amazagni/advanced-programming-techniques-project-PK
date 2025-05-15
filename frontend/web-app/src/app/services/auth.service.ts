import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

export type UserRole = 'guest' | 'user' | 'admin';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private currentRole: UserRole = 'guest';
  private currentUser: string | null = null;

  constructor(private router: Router) {
    const saved = localStorage.getItem('currentUser');
    if (saved) {
      const data = JSON.parse(saved);
      this.currentRole = data.role;
      this.currentUser = data.username;
    }
  }

  login(username: string, password: string): boolean {
    if (username === 'admin' && password === 'admin') {
      this.currentRole = 'admin';
      this.currentUser = 'admin';
    } else {
      // prosty lokalny login usera (każdy login+hasło ok)
      this.currentRole = 'user';
      this.currentUser = username;
    }

    localStorage.setItem('currentUser', JSON.stringify({ username: this.currentUser, role: this.currentRole }));
    return true;
  }

  logout() {
    this.currentRole = 'guest';
    this.currentUser = null;
    localStorage.removeItem('currentUser');
    this.router.navigate(['/']);
  }

  getRole(): UserRole {
    return this.currentRole;
  }

  getUsername(): string | null {
    return this.currentUser;
  }

  isAuthenticated(): boolean {
    return this.currentRole !== 'guest';
  }

  isAdmin(): boolean {
    return this.currentRole === 'admin';
  }
}
