from dataclasses import dataclass
from datetime import datetime
from typing import Optional

@dataclass
class UserSession:
    session_id: str
    user_id: int
    expires_at: int
    created_at: float = datetime.now().timestamp()
    metadata: Optional[dict] = None

    def is_expired(self) -> bool:
        return datetime.now().timestamp() > self.expires_at

    def time_remaining(self) -> float:
        return max(0, self.expires_at - datetime.now().timestamp())

    def update_metadata(self, key: str, value: any) -> None:
        if self.metadata is None:
            self.metadata = {}
        self.metadata[key] = value