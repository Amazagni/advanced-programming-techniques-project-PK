import {Component} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {Item} from '../../models/item.model';
import {HttpClient} from '@angular/common/http';
import Swal from 'sweetalert2';


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

  constructor(
    private fb: FormBuilder,
    private http: HttpClient,
  ) {
    this.itemForm = this.fb.group({
      name: ['', Validators.required],
      description: ['', Validators.required],
      quantity: [1, [Validators.required, Validators.min(1)]],
      image: ['', Validators.required]
    });
  }

  submitForm() {
    if (this.itemForm.valid && this.selectedImage) {
      const formData = new FormData();

      formData.append('name', this.itemForm.get('name')!.value);
      formData.append('description', this.itemForm.get('description')!.value);
      formData.append('quantity', this.itemForm.get('quantity')!.value);
      formData.append('image', this.selectedImage, this.selectedImage.name);

      console.log('FormData contents:');
      for (const [key, value] of (formData as any).entries()) {
        console.log(`${key}: ${value}`);
      }

      this.http.post<Item>('http://localhost:50001/items/create', formData)
        .subscribe({
          next: (response) => {
            console.log('Item added successfully:', response);

            this.itemForm.reset({
              name: '',
              description: '',
              quantity: 1,
              image: ''
            });
            this.selectedImage = null;

            Swal.fire({
              icon: 'success',
              title: 'Sukces!',
              text: 'Przedmiot został dodany poprawnie.',
              confirmButtonText: 'OK',
              background: '#f0f4f8',
              color: '#32446c',
              confirmButtonColor: '#32446c'
            });

          },
          error: (error) => {
            console.error('Error adding item:', error);
            Swal.fire({
              icon: 'error',
              title: 'Błąd!',
              text: 'Wystąpił problem podczas dodawania przedmiotu.',
              confirmButtonText: 'OK',
              background: '#f0f4f8',
              color: '#32446c',
              confirmButtonColor: '#32446c'
            });
          }
        });
    } else {
      console.log('Form is invalid or no image selected.');
    }
  }

  onImageSelected(event: Event) {
    const fileInput = event.target as HTMLInputElement;
    if (fileInput.files && fileInput.files.length > 0) {
      this.selectedImage = fileInput.files[0];
      console.log('Wybrany obrazek:', this.selectedImage);
      this.itemForm.patchValue({ image: this.selectedImage.name });
    }
  }
}
