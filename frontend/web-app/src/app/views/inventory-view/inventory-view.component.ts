import {Component, OnInit} from '@angular/core';
import {ItemCardComponent} from '../../components/item-card/item-card.component';
import {NgForOf, AsyncPipe, NgIf} from '@angular/common';
import {Item} from '../../models/item.model';
import {HttpClient} from '@angular/common/http';
import {map, Observable} from 'rxjs';

@Component({
  selector: 'app-inventory-view',
  standalone: true,
  imports: [
    ItemCardComponent,
    NgForOf,
    AsyncPipe,
    NgIf
  ],
  templateUrl: './inventory-view.component.html',
  styleUrl: './inventory-view.component.css'
})
export class InventoryViewComponent{
  inventory$: Observable<Item[]>;

  constructor(private http: HttpClient) {
    this.inventory$ = this.http.get<Item[]>(
      'http://localhost:50001/items?limit=10&offset=0'
    );
  }

  onQuantityChanged(updatedItem: Item) {
    this.inventory$ = this.inventory$.pipe(
      map(items =>
        items.map(item => (item.Id === updatedItem.Id ? updatedItem : item))
      )
    );
  }

  trackByItemId(index: number, item: Item): number {
    return item.Id;
  }
}
  // [
  //   // id has to be unique !!!
  //   {id: 1, imageUrl: 'assets/images/placeholder.jpg', name: 'Laptop', description: 'Laptop do pracy', quantity: 5},
  //   {id: 2, imageUrl: 'assets/images/placeholder.jpg', name: 'Monitor', description: 'Laptop do pracy', quantity: 5},
  //   {id: 3, imageUrl: 'assets/images/placeholder.jpg', name: 'Klawiatura', description: 'Laptop do pracy', quantity: 5},
  //   {id: 4, imageUrl: 'assets/images/placeholder.jpg', name: 'Laptop', description: 'Laptop do pracy', quantity: 5},
  //   {id: 5, imageUrl: 'assets/images/placeholder.jpg', name: 'Monitor', description: 'Monitor 24 cale', quantity: 3},
  //   {id: 6, imageUrl: 'assets/images/placeholder.jpg', name: 'Klawiatura', description: 'Klawiatura mechaniczna', quantity: 10}
  // ];
