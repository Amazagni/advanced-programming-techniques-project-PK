import {Component, Input} from '@angular/core';
import {ItemCardComponent} from '../../components/item-card/item-card.component';
import {NgForOf} from '@angular/common';
import {Item} from '../../models/item.model';

@Component({
  selector: 'app-inventory-view',
  standalone: true,
  imports: [
    ItemCardComponent,
    NgForOf
  ],
  templateUrl: './inventory-view.component.html',
  styleUrl: './inventory-view.component.css'
})
export class InventoryViewComponent {

  // TODO replace this mock with some real data from backend ;)
  inventory: Item[] = [
    { imageUrl: 'assets/images/placeholder.jpg', name: 'Laptop', description: 'Laptop do pracy', quantity: 5 },
    { imageUrl: 'assets/images/placeholder.jpg', name: 'Monitor', description: 'Laptop do pracy', quantity: 5 },
    { imageUrl: 'assets/images/placeholder.jpg', name: 'Klawiatura', description: 'Laptop do pracy', quantity: 5 },
    { imageUrl: 'assets/images/placeholder.jpg', name: 'Laptop', description: 'Laptop do pracy', quantity: 5 },
    { imageUrl: 'assets/images/placeholder.jpg', name: 'Monitor', description: 'Monitor 24 cale', quantity: 3 },
    { imageUrl: 'assets/images/placeholder.jpg', name: 'Klawiatura', description: 'Klawiatura mechaniczna', quantity: 10 }
  ];
}
