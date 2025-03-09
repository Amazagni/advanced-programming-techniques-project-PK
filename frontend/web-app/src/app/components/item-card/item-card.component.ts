import {Component, Input} from '@angular/core';
import {Item} from '../../models/item.model';

@Component({
  selector: 'app-item-card',
  standalone: true,
  imports: [],
  templateUrl: './item-card.component.html',
  styleUrl: './item-card.component.css'
})
export class ItemCardComponent {
  @Input() item!: Item;

  constructor() {
  }

  openQuantityDialog() {
    console.log("QUANTITY CHANGED")
  }
}
