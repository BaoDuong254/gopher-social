interface PostComment {
  id: number;
  post_id: number;
  user_id: number;
  content: string;
  created_at: string;
  user?: {
    id: number;
    username: string;
  };
}

export interface FeedPost {
  id: number;
  user_id: number;
  comments_count: number;
  content: string;
  created_at: string;
  tags: string[];
  title?: string;
  comments?: PostComment[];
}

interface PostProps {
  post: FeedPost;
  onClick: () => void;
  index?: number;
}

const formatPostDate = (value: string) =>
  new Date(value).toLocaleDateString(undefined, {
    day: "2-digit",
    month: "short",
    year: "numeric",
  });

export const Post: React.FC<PostProps> = ({ post, onClick, index = 0 }) => {
  const tags = post.tags?.filter(Boolean) ?? [];
  const delayClass = `post-delay-${Math.min(index, 8)}`;

  return (
    <article
      className={`post ${delayClass}`}
      onClick={onClick}
      role="button"
      tabIndex={0}
      onKeyDown={(event) => {
        if (event.key === "Enter" || event.key === " ") {
          event.preventDefault();
          onClick();
        }
      }}
    >
      <div className="post-header">
        <h4 className="post-title">{post.title || "Untitled post"}</h4>
        <time className="post-date">{formatPostDate(post.created_at)}</time>
      </div>

      <p className="post-content">{post.content}</p>

      <div className="post-tags">
        {tags.length > 0 ? (
          tags.map((tag) => (
            <span className="tag-chip" key={`${post.id}-${tag}`}>
              {tag}
            </span>
          ))
        ) : (
          <span className="tag-chip tag-chip-muted">general</span>
        )}
      </div>

      <div className="post-bottom">
        <p>
          {post.comments_count} comment{post.comments_count === 1 ? "" : "s"}
        </p>
        <span className="post-cta">Open discussion</span>
      </div>
    </article>
  );
};
