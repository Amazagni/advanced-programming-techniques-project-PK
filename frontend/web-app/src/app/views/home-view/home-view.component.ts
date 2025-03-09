import { Component } from '@angular/core';
import {RouterLink} from '@angular/router';

@Component({
  selector: 'app-home-view',
  standalone: true,
  imports: [
    RouterLink
  ],
  templateUrl: './home-view.component.html',
  styleUrl: './home-view.component.css'
})
export class HomeViewComponent {

  ngOnInit() {
    console.log("HOME VIEW")
  }
}
