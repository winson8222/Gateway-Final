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

jest.mock('Generation Form submisson with different port numbers', () => () => <div>Mock DropdownComponent</div>);
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

describe("Control Panel at different states", () => {
    let getItemSpy: jest.SpyInstance;
  
    beforeEach(() => {
      getItemSpy = jest.spyOn(Storage.prototype, 'getItem');
    });
  
    afterEach(() => {
      getItemSpy.mockRestore();
    });
  
    it("displays control panel when 'generated' is set to 'y' in localStorage", () => {
      getItemSpy.mockImplementation((key) => key === 'generated' ? 'y' : null);
      
      const { queryByText } = render(<SelectServiceComponent />);
      expect(queryByText('Control Panel')).not.toBeNull();
    });
  
    it("does not display control panel when 'generated' is not set to 'y' in localStorage", () => {
      getItemSpy.mockImplementation((key) => 'n');
      
      const { queryByText } = render(<SelectServiceComponent />);
      expect(queryByText('Control Panel')).toBeNull();
    });
  });
  

  describe("SelectServiceComponent", () => {
    let getItemSpy: jest.SpyInstance;
  
    beforeEach(() => {
      getItemSpy = jest.spyOn(Storage.prototype, 'getItem');
    });
  
    afterEach(() => {
      getItemSpy.mockRestore();
    });
  
    it("only Update and Stop buttons are enabled when both 'running' and 'generated' are set to 'y' in localStorage", () => {
      getItemSpy.mockImplementation((key) => 'y');
  
      render(<SelectServiceComponent />);
  
      const updateButton = screen.getByRole('button', { name: /Update|Updating/i });
      const stopButton = screen.getByRole('button', { name: /Stop|Stopping/i });
      const startButton = screen.getByRole('button', { name: /Start|Starting/i });
  
      expect(window.getComputedStyle(updateButton).cursor).toBe('pointer');
      expect(window.getComputedStyle(updateButton).opacity).toBe('1');
  
      expect(window.getComputedStyle(stopButton).cursor).toBe('pointer');
      expect(window.getComputedStyle(stopButton).opacity).toBe('1');
  
      expect(window.getComputedStyle(startButton).cursor).toBe('not-allowed');
      expect(window.getComputedStyle(startButton).opacity).toBe('0.5');
    });
  
    it("only Start button is enabled when both 'running' is set to 'n' and 'generated' is set to 'y' in localStorage", () => {
      getItemSpy.mockImplementation((key) => {
        if (key === 'running') return 'n';
        if (key === 'generated') return 'y';
        return null;
      });
  
      render(<SelectServiceComponent />);
      
      const updateButton = screen.getByRole('button', { name: /Update|Updating/i });
      const stopButton = screen.getByRole('button', { name: /Stop|Stopping/i });
      const startButton = screen.getByRole('button', { name: /Start|Starting/i });
  
      expect(window.getComputedStyle(updateButton).cursor).toBe('not-allowed');
      expect(window.getComputedStyle(updateButton).opacity).toBe('0.5');
  
      expect(window.getComputedStyle(stopButton).cursor).toBe('not-allowed');
      expect(window.getComputedStyle(stopButton).opacity).toBe('0.5');
  
      expect(window.getComputedStyle(startButton).cursor).toBe('pointer');
      expect(window.getComputedStyle(startButton).opacity).toBe('1');
    });
  });
  



