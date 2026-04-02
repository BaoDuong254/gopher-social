import { useNavigate, useParams } from "react-router-dom";

function ConfirmPage() {
  const { token = "" } = useParams();
  const redirect = useNavigate();
  const handleConfirm = async () => {
    const response = await fetch(
      `${import.meta.env.VITE_API_URL}/users/activate/${token}`,
      {
        method: "PUT",
      },
    );
    if (response.ok) {
      redirect("/");
    } else {
      alert("Activation failed. Please try again.");
    }
  };
  return (
    <div className="confirm-page">
      <h1>Confirm Your Account</h1>
      <p>Please click the button below to confirm your account.</p>
      <button onClick={handleConfirm}>Confirm Account</button>
    </div>
  );
}
export default ConfirmPage;
