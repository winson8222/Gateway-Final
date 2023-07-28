import { render, fireEvent, screen, waitFor } from '@testing-library/react';
import React from 'react';
import SelectServiceComponent from '../components/SelectService'; // Update this import to match your actual file structure
import axios ,{ AxiosResponse } from 'axios';
import { toast, ToastContent, ToastOptions, Id } from 'react-toastify';
import '@testing-library/jest-dom'


jest.mock('axios');
jest.mock('react-toastify', () => ({
  toast: {
    error: jest.fn(),
    success: jest.fn(),
    isActive: jest.fn(),
  }
}));

jest.mock('../components/MicrosvcDropdown', () => () => <div>Mock DropdownComponent</div>);
describe('SelectServiceComponent', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it('prevents form submission when url input is blank', async () => {
    render(<SelectServiceComponent />);

    // Find the url input field and the submit button
    const urlInput = screen.getByLabelText('Gateway Port');
    const submitButton = screen.getByText('Generate');

    // Attempt to submit the form with empty url input
    fireEvent.change(urlInput, { target: { value: '' } });
    fireEvent.click(submitButton);

    // Wait for the toast error to be called
    await waitFor(() => {
      expect(toast.error).toHaveBeenCalledWith('Fields not filled');
    });
  });

  it('prevents form submission when url input is 8888, 8889, 8890 or 20000', async () => {
    render(<SelectServiceComponent />);

    // Define an array of invalid inputs
    const invalidInputs = ['8888', '8889', '8890', '20000'];

    const urlInput = screen.getByLabelText('Gateway Port');
    const submitButton = screen.getByText('Generate');

    // Iterate over the invalid inputs and test each one
    for (let invalidInput of invalidInputs) {
      fireEvent.change(urlInput, { target: { value: invalidInput } });
      fireEvent.click(submitButton);

      await waitFor(() => {
        expect(toast.error).toHaveBeenCalledWith('Ports 8888, 8889, 8890, 20000 are not allowed');
      });
    }
  });
});

describe("CreateServiceComponent", () => {
    beforeEach(() => {
      jest.clearAllMocks();
    });
  
    it("sends axios request to /gen when form condition is fulfilled", async () => {
      // Mock the toast.error function to avoid displaying error messages during the test
      const mockToastError = jest.fn();
      jest.spyOn(toast, 'error').mockImplementation(mockToastError);
  
      // Mock the localStorage.getItem method to return "n" for both "generated" and "running" keys
      const mockLocalStorage = {
        getItem: jest.fn((key) => (key === "generated" || key === "running") ? "n" : null),
      };
      Object.defineProperty(window, "localStorage", { value: mockLocalStorage });
  
      const formData = { url: "1234", lb: "ROUND_ROBIN" };
      const mockResponse: Partial<AxiosResponse> = { data: { outcome: "Gateway Generated" }, status: 200};
  
      // Mock the axios.post function to return a custom response
      jest.spyOn(axios, 'post').mockResolvedValue(mockResponse);
  
      render(<SelectServiceComponent />);
  
      // Fill up the form inputs
      fireEvent.change(screen.getByLabelText("Gateway Port"), { target: { value: "1234" } });
      fireEvent.change(screen.getByLabelText("Load Balancing Type"), { target: { value: "ROUND_ROBIN" } });
  
      // Attempt to submit the form
      fireEvent.click(screen.getByText("Generate"));
  
      // Wait for the axios request to be resolved
      await waitFor(() => {
        expect(axios.post).toHaveBeenCalledWith("http://localhost:3333/gen", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
        expect(mockToastError).not.toHaveBeenCalled();
      });
    });
  });
  
  
  
  
  
  
  

