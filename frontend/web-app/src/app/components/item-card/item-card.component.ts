import {Component, Input, Output, EventEmitter} from '@angular/core';
import {Item} from '../../models/item.model';
import {FormsModule} from '@angular/forms';
import {HttpClient} from '@angular/common/http';

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
  @Output() quantityChanged = new EventEmitter<Item>();
  newQuantity: number | undefined;
  isDialogOpen = false;

  constructor(private http: HttpClient) {}

  openQuantityDialog() {
    this.newQuantity = this.item.Quantity;
    const itemId = this.item.Id;
    if (itemId) {
      const modalElement = document.getElementById(itemId.toString());
      if (modalElement) {
        const modal = new bootstrap.Modal(modalElement);
        modal.show();
        this.isDialogOpen = true;
        return;
      }
    }
    console.error('Quantity modal not found');
  }

  saveQuantityChange() {
    const itemId = this.item.Id;
    if (itemId && this.newQuantity !== undefined) {
      this.http.put<Item>(`http://localhost:50001/items/update/${itemId}`, 
        { Quantity: this.newQuantity })
        .subscribe({
          next: (updatedItem) => {
            console.log('Quantity updated successfully:', updatedItem);
            this.quantityChanged.emit(updatedItem); // Emit the updated item
            const modalElement = document.getElementById(itemId.toString());
            if (modalElement && this.isDialogOpen) {
              const modal = bootstrap.Modal.getInstance(modalElement);
              modal.hide();
              this.isDialogOpen = false;
            }
          },
          error: (error) => {
            console.error('Error updating quantity:', error);
          }
        });
    }
  }
}