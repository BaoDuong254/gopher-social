import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { API_URL } from "./App";

export const ConfirmationPage = () => {
  const { token = "" } = useParams();
  const redirect = useNavigate();
  const [statusMessage, setStatusMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleConfirm = async () => {
    setStatusMessage("");
    setIsSubmitting(true);

    try {
      const response = await fetch(`${API_URL}/users/activate/${token}`, {
        method: "PUT",
      });

      if (response.ok) {
        redirect("/");
        return;
      }

      setStatusMessage(
        `Activation failed (${response.status}). Please check the link and try again.`,
      );
    } catch {
      setStatusMessage(
        "Unable to activate your account right now. Please try again later.",
      );
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="page-shell">
      <main className="confirmation-layout">
        <section className="status-card confirmation-card">
          <p className="eyebrow">Account activation</p>
          <h1>Confirm your invitation</h1>
          <p>
            Finish account setup to start posting, following, and engaging with
            your gopher network.
          </p>

          {statusMessage && <p className="form-alert">{statusMessage}</p>}

          <button
            className="btn btn-primary"
            onClick={handleConfirm}
            disabled={isSubmitting}
          >
            {isSubmitting ? "Confirming..." : "Activate account"}
          </button>
        </section>
      </main>
    </div>
  );
};
