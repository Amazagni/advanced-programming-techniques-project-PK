import {Component, Input} from '@angular/core';
import {Item} from '../../models/item.model';
import {FormsModule} from '@angular/forms';

@Component({
  selector: 'app-item-card',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './item-card.component.html',
  styleUrl: './item-card.component.css'
})
export class ItemCardComponent {
  @Input() item!: Item;

  constructor() {
  }

  openQuantityDialog() {
    const itemId = this.item.id;
    console.log("JEST", itemId)
    if (itemId) {
      console.log("ITEM ID")
      const modalElement = document.getElementById(itemId.toString());
      if (modalElement) {
        const modal = new bootstrap.Modal(modalElement);
        modal.show();
        return;
      }
    }
    console.error('Quantity modal not found');
  }

  saveQuantityChange() {
    const itemId = this.item.id;
    if (itemId) {
      const modalElement = document.getElementById(itemId.toString());
      if (modalElement) {
        // TODO UPDATE QUANTITY IN DATABASE (this.item is a updated version of item)
        const modal = bootstrap.Modal.getInstance(modalElement);
        modal.hide();
      }
    }
  }
}
