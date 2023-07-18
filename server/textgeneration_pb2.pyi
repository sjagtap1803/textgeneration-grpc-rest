from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class FileUpload(_message.Message):
    __slots__ = ["filedata", "filename"]
    FILEDATA_FIELD_NUMBER: _ClassVar[int]
    FILENAME_FIELD_NUMBER: _ClassVar[int]
    filedata: bytes
    filename: str
    def __init__(self, filename: _Optional[str] = ..., filedata: _Optional[bytes] = ...) -> None: ...

class TextGenPrompt(_message.Message):
    __slots__ = ["cleanup", "prompt"]
    CLEANUP_FIELD_NUMBER: _ClassVar[int]
    PROMPT_FIELD_NUMBER: _ClassVar[int]
    cleanup: bool
    prompt: str
    def __init__(self, prompt: _Optional[str] = ..., cleanup: bool = ...) -> None: ...

class TextGenResponse(_message.Message):
    __slots__ = ["text"]
    TEXT_FIELD_NUMBER: _ClassVar[int]
    text: str
    def __init__(self, text: _Optional[str] = ...) -> None: ...
