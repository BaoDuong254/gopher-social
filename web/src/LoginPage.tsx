import { useState } from "react";
import { API_URL } from "./App";
import { useNavigate } from "react-router-dom";
import { useCookies } from "react-cookie";

type LoginResponse = {
  data?: string;
  error?: string;
};

export const LoginPage: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const redirect = useNavigate();
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [_, setCookie] = useCookies(["at"]);

  const handleLogin = async () => {
    setErrorMessage("");
    setIsSubmitting(true);

    try {
      const response = await fetch(`${API_URL}/authentication/token`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      const raw = await response.text();

      let out: LoginResponse = {};
      if (raw.length > 0) {
        try {
          out = JSON.parse(raw);
        } catch {
          throw new Error(
            `Server returned invalid response (status ${response.status})`,
          );
        }
      }

      if (!response.ok) {
        throw new Error(
          out.error || `Login failed with status ${response.status}`,
        );
      }

      if (!out.data) {
        throw new Error("Login succeeded but token was missing");
      }

      setCookie("at", out.data, { path: "/" });

      console.log(out);

      redirect("/feed");
    } catch (error) {
      setErrorMessage(
        error instanceof Error ? error.message : "Unexpected login error",
      );
      console.log({ error });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="page-shell">
      <div className="login-wrap">
        <section className="login-panel">
          <p className="eyebrow">Welcome back</p>
          <h1>Sign in to GopherSocial</h1>
          <p className="login-subtitle">
            Stay close to your dev community, publish updates, and follow the
            conversations that matter.
          </p>

          <label className="field-group">
            <span className="field-label">Email</span>
            <input
              className="text-field"
              type="email"
              placeholder="you@example.com"
              value={email}
              onChange={(v) => setEmail(v.target.value)}
            />
          </label>

          <label className="field-group">
            <span className="field-label">Password</span>
            <input
              className="text-field"
              type="password"
              placeholder="Enter your password"
              value={password}
              onChange={(v) => setPassword(v.target.value)}
            />
          </label>

          {errorMessage && (
            <p className="form-alert" role="alert">
              {errorMessage}
            </p>
          )}

          <button
            className="btn btn-primary login-button"
            onClick={handleLogin}
            disabled={isSubmitting}
          >
            {isSubmitting ? "Logging in..." : "Log in"}
          </button>
        </section>

        <aside className="login-side-card">
          <h2>What is new?</h2>
          <ul>
            <li>Faster feed loading for active communities.</li>
            <li>Cleaner post layouts for easier reading.</li>
            <li>Keyboard-friendly navigation across cards.</li>
          </ul>
        </aside>
      </div>
    </div>
  );
};
