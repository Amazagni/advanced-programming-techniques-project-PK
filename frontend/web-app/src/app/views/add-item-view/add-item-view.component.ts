import { Component } from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Item} from '../../models/item.model';

@Component({
  selector: 'app-add-item-view',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './add-item-view.component.html',
  styleUrl: './add-item-view.component.css'
})
export class AddItemViewComponent {
  itemForm: FormGroup;
  selectedImage: File | null = null;

  constructor(private fb: FormBuilder) {
    this.itemForm = this.fb.group({
      name: ['', Validators.required],
      description: ['', Validators.required],
      quantity: [1, [Validators.required, Validators.min(1)]],
      imageUrl: ['', Validators.required]
    });
  }

  submitForm() {
    if (this.itemForm.valid && this.selectedImage) {
      const newItem: Item = {
        ...this.itemForm.value,
        image: this.selectedImage.name
      };
      console.log('Nowy item:', newItem);
    }
  }

  onImageSelected(event: Event) {
    const fileInput = event.target as HTMLInputElement;
    if (fileInput.files && fileInput.files.length > 0) {
      this.selectedImage = fileInput.files[0];
      console.log('Wybrany obrazek:', this.selectedImage);
    }
  }
}
