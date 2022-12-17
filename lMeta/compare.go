package lMeta

func (meta *Meta) Compare(o *Meta) bool {
	if o == nil {
		return false
	}

	if meta.Size != o.Size {
		return false
	}

	if meta.SHA1 != o.SHA1 {
		return false
	}

	if meta.MD5 != o.MD5 {
		return false
	}

	if meta.ED2K != o.ED2K {
		return false
	}

	if meta.SHA256 != o.SHA256 {
		return false
	}

	return true
}
