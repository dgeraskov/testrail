package testrail

import (
	"fmt"
)

// GetAttachment returns the existing attachment contents by its id
func (c *Client) GetAttachment(attachmentID int) ([]byte, error) {
	attachment := RawData{}
	err := c.sendRequest("GET", fmt.Sprintf("get_attachment/%d", attachmentID), nil, &attachment)
	return attachment.data, err
}
