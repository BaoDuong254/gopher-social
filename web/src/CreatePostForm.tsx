import { useState } from "react";
import { API_URL } from "./App";
import { useCookies } from "react-cookie";

export const CreatePostForm: React.FC<{ onFetchPosts: () => void }> = ({
  onFetchPosts,
}) => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [message, setMessage] = useState("");

  const [cookies] = useCookies(["at"]);
  const at = cookies.at;

  const handleSubmit = async () => {
    if (!title.trim() || !content.trim()) {
      setMessage("Please add a title and content before posting.");
      return;
    }

    setMessage("");
    setIsSubmitting(true);

    try {
      const response = await fetch(`${API_URL}/posts`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${at}`,
        },
        body: JSON.stringify({
          title,
          content,
        }),
      });

      if (!response.ok) {
        throw new Error(`Unable to publish post (${response.status})`);
      }

      setTitle("");
      setContent("");
      setMessage("Posted successfully.");
      onFetchPosts();
    } catch (error) {
      setMessage(
        error instanceof Error
          ? error.message
          : "Unexpected error while posting.",
      );
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="composer-card">
      <label className="field-group">
        <span className="field-label">Title</span>
        <input
          className="text-field"
          placeholder="What are you building today?"
          value={title}
          type="text"
          onChange={(e) => setTitle(e.target.value)}
        />
      </label>
      <label className="field-group">
        <span className="field-label">Content</span>
        <textarea
          className="text-area"
          placeholder="Share progress, ask for feedback, or drop an idea..."
          value={content}
          onChange={(e) => setContent(e.target.value)}
        />
      </label>

      {message && (
        <p className="composer-message" role="status">
          {message}
        </p>
      )}

      <button
        className="btn btn-primary composer-button"
        onClick={handleSubmit}
        disabled={isSubmitting}
      >
        {isSubmitting ? "Publishing..." : "Share post"}
      </button>
    </div>
  );
};
