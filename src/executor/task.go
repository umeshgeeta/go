/*
 * MIT License
 * Author: Umesh Patil, Neosemantix, Inc.
 */
package executor

type Task interface {
	GetId() int

	Execute() Response

	SetRespChan(rc chan Response)

	GetRespChan() chan Response

	IsBlocking() bool
}
