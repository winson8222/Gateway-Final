import { render, fireEvent, screen, waitFor } from "@testing-library/react";
import React from 'react';
import CreateServiceComponent from "../components/CreateService";
import { toast } from 'react-toastify';
import { useRouter } from 'next/router';

// ... your tests


jest.mock('react-toastify', () => ({
  toast: {
    error: jest.fn(),
    success: jest.fn(),
    isActive: jest.fn(),
  }
}));

jest.mock('next/router');


describe("CreateServiceComponent", () => {
  it("prevents form submission when any input is blank", async () => {
    render(<CreateServiceComponent />);

    // Attempt to submit the form with empty inputs
    fireEvent.click(screen.getByText("Submit"));

    // Wait for the toast error to be called
    await waitFor(() => {
      expect(toast.error).toHaveBeenCalledWith('Please fill up all columns.');
    });
  });
});
